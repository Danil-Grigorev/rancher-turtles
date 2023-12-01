/*
Copyright SUSE 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"fmt"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	"sigs.k8s.io/cluster-api/util/conditions"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/source"

	operatorv1 "sigs.k8s.io/cluster-api-operator/api/v1alpha2"

	turtlesv1 "github.com/rancher-sandbox/rancher-turtles/api/v1alpha1"
)

const fieldOwner = "capi-provider-operator"

// CAPIProviderReconciler reconciles a CAPIProvider object.
type CAPIProviderReconciler struct {
	Client client.WithWatch
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=turtles-capi.cattle.io,resources=capiproviders,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=turtles-capi.cattle.io,resources=capiproviders/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=turtles-capi.cattle.io,resources=capiproviders/finalizers,verbs=update
//+kubebuilder:rbac:groups=operator.cluster.x-k8s.io,resources=*,verbs=get;list;watch;create;update;patch;delete

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the CAPIProvider object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.16.3/pkg/reconcile
func (r *CAPIProviderReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)

	capiProvider := &turtlesv1.CAPIProvider{ObjectMeta: metav1.ObjectMeta{
		Namespace: req.Namespace,
		Name:      req.Name,
	}}

	if err := r.Client.Get(ctx, client.ObjectKeyFromObject(capiProvider), capiProvider); client.IgnoreNotFound(err) != nil {
		log.Error(err, fmt.Sprintf("Unable to get CAPIProvider manifest: %s", req.String()))

		return ctrl.Result{}, err
	}

	return r.reconcileNormal(ctx, capiProvider)
}

func (r *CAPIProviderReconciler) reconcileNormal(ctx context.Context, capiProvider *turtlesv1.CAPIProvider) (_ ctrl.Result, err error) {
	return r.sync(ctx, capiProvider)
}

func (r *CAPIProviderReconciler) sync(ctx context.Context, capiProvider *turtlesv1.CAPIProvider) (_ ctrl.Result, err error) {
	log := log.FromContext(ctx)

	mirror := NewMirror(capiProvider)
	cl := NewMirrorClient(r.Client, mirror)

	if err := cl.Get(ctx, client.ObjectKeyFromObject(capiProvider), capiProvider); client.IgnoreNotFound(err) != nil {
		log.Error(err, fmt.Sprintf("Unable to get mirrored manifest: %s", client.ObjectKeyFromObject(capiProvider).String()))

		return ctrl.Result{}, err
	}

	defer patcher(ctx, cl, capiProvider, &err)

	mirror.Sync()

	switch {
	case conditions.IsTrue(capiProvider, operatorv1.ProviderInstalledCondition):
		capiProvider.Status.Phase = turtlesv1.Ready
	case conditions.IsFalse(capiProvider, operatorv1.PreflightCheckCondition):
		capiProvider.Status.Phase = turtlesv1.Failed
	default:
		capiProvider.Status.Phase = turtlesv1.Provisioning
	}

	return ctrl.Result{}, nil
}

func patcher(ctx context.Context, c client.Client, obj client.Object, reterr *error) {
	// Always attempt to update the object and status after each reconciliation.
	log := log.FromContext(ctx)
	log.Info(fmt.Sprintf("Updating object %s", client.ObjectKeyFromObject(obj)))

	patchOptions := []client.PatchOption{
		client.ForceOwnership,
		client.FieldOwner(fieldOwner),
	}
	statusOptions := []client.SubResourcePatchOption{
		client.ForceOwnership,
		client.FieldOwner(fieldOwner),
	}

	obj.SetManagedFields(nil)

	if err := c.Status().Patch(ctx, obj, client.Apply, statusOptions...); err != nil {
		if apierrors.IsNotFound(err) {
			log.Info(fmt.Sprintf("Object %s is not found, skipping update", client.ObjectKeyFromObject(obj)))
			return
		}

		*reterr = kerrors.NewAggregate([]error{*reterr, err})
		log.Error(*reterr, fmt.Sprintf("Unable to patch object status: %s", *reterr))
	}

	obj.SetManagedFields(nil)

	if err := c.Patch(ctx, obj, client.Apply, patchOptions...); err != nil {
		if apierrors.IsNotFound(err) {
			log.Info(fmt.Sprintf("Object %s is not found, skipping update", client.ObjectKeyFromObject(obj)))
			return
		}

		*reterr = kerrors.NewAggregate([]error{*reterr, err})
		log.Error(*reterr, fmt.Sprintf("Unable to patch object: %s", *reterr))
	}
}

// SetupWithManager sets up the controller with the Manager.
func (r *CAPIProviderReconciler) SetupWithManager(ctx context.Context, mgr ctrl.Manager) error {
	_ = log.FromContext(ctx)

	c, err := ctrl.NewControllerManagedBy(mgr).
		For(&turtlesv1.CAPIProvider{}).
		Build(r)
	if err != nil {
		return fmt.Errorf("creating new CAPI Provider controller: %w", err)
	}

	resources := []client.Object{
		&operatorv1.CoreProvider{},
		&operatorv1.ControlPlaneProvider{},
		&operatorv1.InfrastructureProvider{},
		&operatorv1.BootstrapProvider{},
		&operatorv1.AddonProvider{},
	}

	for _, resource := range resources {
		err = c.Watch(
			source.Kind(mgr.GetCache(), resource),
			handler.EnqueueRequestForOwner(mgr.GetScheme(), mgr.GetRESTMapper(), &turtlesv1.CAPIProvider{}),
			predicate.ResourceVersionChangedPredicate{},
		)
		if err != nil {
			return fmt.Errorf("creating new provider watches: %w", err)
		}
	}

	return nil
}
