package sync

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	turtlesv1 "github.com/rancher-sandbox/rancher-turtles/api/v1alpha1"
)

// SecretSyncer is a structure mirroring variable secret state of the CAPI Operator Provider object.
type SecretSyncer struct {
	*DefaultSyncer

	Secret *corev1.Secret
}

// NewSecretSyncer creates a new secret object sync.
func NewSecretSyncer(cl client.Client, capiProvider *turtlesv1.CAPIProvider) Syncer {
	secret := SecretSyncer{}.GetSecret(capiProvider)

	return &SecretSyncer{
		DefaultSyncer: NewDefaultSyncer(cl, capiProvider, secret),
		Secret:        secret,
	}
}

// GetSecret returning the mirrored secret resource template.
func (SecretSyncer) GetSecret(capiProvider *turtlesv1.CAPIProvider) *corev1.Secret {
	meta := metav1.ObjectMeta{
		Name:      capiProvider.Name,
		Namespace: capiProvider.Namespace,
	}

	if capiProvider.Spec.ConfigSecret != nil {
		meta.Name = capiProvider.Spec.ConfigSecret.Name
	}

	return &corev1.Secret{ObjectMeta: meta}
}

// Template returning the mirrored secret resource template.
func (SecretSyncer) Template(capiProvider *turtlesv1.CAPIProvider) client.Object {
	return SecretSyncer{}.GetSecret(capiProvider)
}

// Sync updates the mirror object state from the upstream source object
// Direction of updates:
// Spec -> down
// up <- Status.
func (s *SecretSyncer) Sync(_ context.Context) error {
	s.SyncObjects()

	return nil
}

// SyncObjects updates the Source CAPIProvider object and the environment secret state.
// Direction of updates:
// Spec.Features + Spec.Variables -> Status.Variables -> Secret.
func (s *SecretSyncer) SyncObjects() {
	setVariables(s.DefaultSyncer.Source)
	setFeatures(s.DefaultSyncer.Source)

	s.Secret.StringData = s.DefaultSyncer.Source.Status.Variables
}

func setVariables(capiProvider *turtlesv1.CAPIProvider) {
	if capiProvider.Spec.Variables != nil {
		capiProvider.Status.Variables = capiProvider.Spec.Variables
	}
}

func setFeatures(capiProvider *turtlesv1.CAPIProvider) {
	value := "true"
	features := capiProvider.Spec.Features
	variables := capiProvider.Status.Variables

	if features != nil {
		switch {
		case features.ClusterResourceSet:
			variables["EXP_CLUSTER_RESOURCE_SET"] = value
			fallthrough
		case features.ClusterTopology:
			variables["CLUSTER_TOPOLOGY"] = value
			fallthrough
		case features.MachinePool:
			variables["EXP_MACHINE_POOL"] = value
			fallthrough
		default:
		}
	}
}
