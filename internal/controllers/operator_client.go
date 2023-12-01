package controllers

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/interceptor"

	operatorv1 "sigs.k8s.io/cluster-api-operator/api/v1alpha2"

	turtlesv1 "github.com/rancher-sandbox/rancher-turtles/api/v1alpha1"
)

// Mirror is a structure mirroring state of the CAPI Operator Provider object.
type Mirror struct {
	Source *turtlesv1.CAPIProvider
	mirror client.Object
}

// NewMirror creates a new mirror object.
func NewMirror(capiProvider *turtlesv1.CAPIProvider) *Mirror {
	return (&Mirror{
		Source: capiProvider,
		mirror: capiProvider.GetCAPIProvider(),
	}).Sync()
}

func (m *Mirror) setOwnerReference(obj client.Object) {
	if m.Source.GetUID() != "" {
		obj.SetFinalizers([]string{metav1.FinalizerDeleteDependents})
		obj.SetOwnerReferences([]metav1.OwnerReference{{
			APIVersion:         turtlesv1.GroupVersion.String(),
			Kind:               turtlesv1.ProviderKind,
			Name:               m.Source.GetName(),
			UID:                m.Source.GetUID(),
			Controller:         pointer.Bool(true),
			BlockOwnerDeletion: pointer.Bool(true),
		}})
	}
}

// Sync updates the mirror object state from the upstream source object
// Direcrion of updates:
// Spec -> down
// up <- Status.
func (m *Mirror) Sync() *Mirror {
	m.Source.APIVersion = turtlesv1.GroupVersion.String()
	m.Source.Kind = turtlesv1.ProviderKind

	dst := m.Source.GetCAPIProvider()
	m.setOwnerReference(dst)

	switch mirror := m.mirror.(type) {
	case *operatorv1.InfrastructureProvider:
		dst, matches := dst.(*operatorv1.InfrastructureProvider)
		if matches {
			m.Source.Spec.ProviderSpec.DeepCopyInto(&dst.Spec.ProviderSpec)
			mirror.Status.ProviderStatus.DeepCopyInto(&m.Source.Status.ProviderStatus)
		}
	case *operatorv1.CoreProvider:
		dst, matches := dst.(*operatorv1.CoreProvider)
		if matches {
			m.Source.Spec.ProviderSpec.DeepCopyInto(&dst.Spec.ProviderSpec)
			mirror.Status.ProviderStatus.DeepCopyInto(&m.Source.Status.ProviderStatus)
		}
	case *operatorv1.ControlPlaneProvider:
		dst, matches := dst.(*operatorv1.ControlPlaneProvider)
		if matches {
			m.Source.Spec.ProviderSpec.DeepCopyInto(&dst.Spec.ProviderSpec)
			mirror.Status.ProviderStatus.DeepCopyInto(&m.Source.Status.ProviderStatus)
		}
	case *operatorv1.BootstrapProvider:
		dst, matches := dst.(*operatorv1.BootstrapProvider)
		if matches {
			m.Source.Spec.ProviderSpec.DeepCopyInto(&dst.Spec.ProviderSpec)
			mirror.Status.ProviderStatus.DeepCopyInto(&m.Source.Status.ProviderStatus)
		}
	case *operatorv1.AddonProvider:
		dst, matches := dst.(*operatorv1.AddonProvider)
		if matches {
			m.Source.Spec.ProviderSpec.DeepCopyInto(&dst.Spec.ProviderSpec)
			mirror.Status.ProviderStatus.DeepCopyInto(&m.Source.Status.ProviderStatus)
		}
	default:
	}

	m.mirror = dst

	return m
}

// Get will collect both source and mirror objects from the cluster.
func (m *Mirror) Get(ctx context.Context, cl client.WithWatch, key client.ObjectKey, obj client.Object, opts ...client.GetOption) error {
	if err := cl.Get(ctx, key, obj, opts...); err != nil {
		return err
	}

	return cl.Get(ctx, client.ObjectKeyFromObject(m.mirror), m.mirror, opts...)
}

// Patch will only patch mirror object in the cluster.
func (m *Mirror) Patch(ctx context.Context, cl client.WithWatch, _ client.Object, patch client.Patch, opts ...client.PatchOption) error {
	return cl.Patch(ctx, m.mirror, patch, opts...)
}

func (m *Mirror) funcs() interceptor.Funcs {
	return interceptor.Funcs{
		Get:   m.Get,
		Patch: m.Patch,
	}
}

// NewMirrorClient is an interceptor client, keeping mirrored object in sync on request.
func NewMirrorClient(interceptedClient client.WithWatch, provider *Mirror) client.WithWatch {
	return interceptor.NewClient(interceptedClient, provider.funcs())
}
