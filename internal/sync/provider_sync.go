package sync

import (
	"context"

	"sigs.k8s.io/controller-runtime/pkg/client"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	operatorv1 "sigs.k8s.io/cluster-api-operator/api/v1alpha2"

	turtlesv1 "github.com/rancher-sandbox/rancher-turtles/api/v1alpha1"
)

// ProviderSyncer is a structure mirroring state of the CAPI Operator Provider object.
type ProviderSyncer struct {
	*DefaultSyncer
}

// NewProviderSyncer creates a new mirror object.
func NewProviderSyncer(cl client.Client, capiProvider *turtlesv1.CAPIProvider) Syncer {
	return &ProviderSyncer{
		DefaultSyncer: NewDefaultSyncer(cl, capiProvider, ProviderSyncer{}.Template(capiProvider)),
	}
}

// Template returning the mirrored CAPI Operator manifest template.
func (ProviderSyncer) Template(capiProvider *turtlesv1.CAPIProvider) client.Object {
	meta := metav1.ObjectMeta{
		Name:      string(capiProvider.Spec.Name),
		Namespace: capiProvider.GetNamespace(),
	}

	switch capiProvider.Spec.Type {
	case turtlesv1.InfrastructureProvider:
		return &operatorv1.InfrastructureProvider{ObjectMeta: meta}
	case turtlesv1.CoreProvider:
		return &operatorv1.CoreProvider{ObjectMeta: meta}
	case turtlesv1.ControlPlaneProvider:
		return &operatorv1.ControlPlaneProvider{ObjectMeta: meta}
	case turtlesv1.BootstrapProvider:
		return &operatorv1.BootstrapProvider{ObjectMeta: meta}
	case turtlesv1.AddonProvider:
		return &operatorv1.AddonProvider{ObjectMeta: meta}
	default:
	}

	return nil
}

// Sync updates the mirror object state from the upstream source object
// Direction of updates:
// Spec -> down
// up <- Status.
func (s *ProviderSyncer) Sync(_ context.Context) error {
	s.SyncObjects()

	return nil
}

// SyncObjects updates the Source CAPIProvider object and the destination provider object states.
// Direction of updates:
// Spec -> <Common>Provider
// CAPIProvider <- Status.
func (s *ProviderSyncer) SyncObjects() {
	switch mirror := s.Destination.(type) {
	case *operatorv1.InfrastructureProvider:
		s.Source.Spec.ProviderSpec.DeepCopyInto(&mirror.Spec.ProviderSpec)
		mirror.Status.ProviderStatus.DeepCopyInto(&s.Source.Status.ProviderStatus)
	case *operatorv1.CoreProvider:
		s.Source.Spec.ProviderSpec.DeepCopyInto(&mirror.Spec.ProviderSpec)
		mirror.Status.ProviderStatus.DeepCopyInto(&s.Source.Status.ProviderStatus)
	case *operatorv1.ControlPlaneProvider:
		s.Source.Spec.ProviderSpec.DeepCopyInto(&mirror.Spec.ProviderSpec)
		mirror.Status.ProviderStatus.DeepCopyInto(&s.Source.Status.ProviderStatus)
	case *operatorv1.BootstrapProvider:
		s.Source.Spec.ProviderSpec.DeepCopyInto(&mirror.Spec.ProviderSpec)
		mirror.Status.ProviderStatus.DeepCopyInto(&s.Source.Status.ProviderStatus)
	case *operatorv1.AddonProvider:
		s.Source.Spec.ProviderSpec.DeepCopyInto(&mirror.Spec.ProviderSpec)
		mirror.Status.ProviderStatus.DeepCopyInto(&s.Source.Status.ProviderStatus)
	default:
	}
}
