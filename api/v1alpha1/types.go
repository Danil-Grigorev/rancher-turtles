package v1alpha1

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	operatorv1 "sigs.k8s.io/cluster-api-operator/api/v1alpha2"
)

// ProviderKind is a CAPIProvider kind string.
const ProviderKind = "CAPIProvider"

// ProviderName defines the designated provider infrastructure provider name.
type ProviderName string

const (
	// AWSProvider is the name for the aws provider.
	AWSProvider ProviderName = "aws"
	// AzureProvider is the name for Azure provider.
	AzureProvider ProviderName = "azure"
	// GCPProvider is the name for the GCP provider.
	GCPProvider ProviderName = "gcp"
	// DockerProvider is the name for the docker provider.
	DockerProvider ProviderName = "docker"
	// RKE2Provider is the name for the RKE2 provider.
	RKE2Provider ProviderName = "rke2"
)

// ProviderType defines the type of the CAPI Provider.
type ProviderType string

const (
	// InfrastructureProvider is the name for the infrastructure CAPI Provider.
	InfrastructureProvider ProviderType = "infrastructure"
	// CoreProvider is the name for core CAPI Provider.
	CoreProvider ProviderType = "core"
	// ControlPlaneProvider is the name for the controlPlane CAPI Provider.
	ControlPlaneProvider ProviderType = "controlPlane"
	// BootstrapProvider is the name for the bootstrap CAPI Provider.
	BootstrapProvider ProviderType = "bootstrap"
	// AddonProvider is the name for the addon CAPI Provider.
	AddonProvider ProviderType = "addon"
)

// ToKind converts ProviderType to CAPI Operator provider object kind.
func (t ProviderType) ToKind() string {
	return cases.Title(language.English).String(string(t)) + "Provider"
}

// GetCAPITypeMeta returning type and metadata for the mirrored CAPI Operator manifest.
func (p *CAPIProvider) GetCAPITypeMeta() *CAPIProvider {
	return &CAPIProvider{
		TypeMeta: metav1.TypeMeta{
			APIVersion: operatorv1.GroupVersion.String(),
			Kind:       p.Spec.Type.ToKind(),
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      string(p.Spec.Name),
			Namespace: p.GetNamespace(),
		},
	}
}

// GetCAPIProvider returning the mirrored CAPI Operator manifest template.
func (p CAPIProvider) GetCAPIProvider() client.Object {
	obj := p.GetCAPITypeMeta()

	switch p.Spec.Type {
	case InfrastructureProvider:
		return &operatorv1.InfrastructureProvider{ObjectMeta: obj.ObjectMeta, TypeMeta: obj.TypeMeta}
	case CoreProvider:
		return &operatorv1.CoreProvider{ObjectMeta: obj.ObjectMeta, TypeMeta: obj.TypeMeta}
	case ControlPlaneProvider:
		return &operatorv1.ControlPlaneProvider{ObjectMeta: obj.ObjectMeta, TypeMeta: obj.TypeMeta}
	case BootstrapProvider:
		return &operatorv1.BootstrapProvider{ObjectMeta: obj.ObjectMeta, TypeMeta: obj.TypeMeta}
	case AddonProvider:
		return &operatorv1.AddonProvider{ObjectMeta: obj.ObjectMeta, TypeMeta: obj.TypeMeta}
	default:
		return nil
	}
}

// Phase defines the current state of the CAPI Provider resource.
type Phase string

const (
	// Pending status identifies a provder which has not yet started provisioning.
	Pending Phase = "Pending"
	// Provisioning status defines provider in a provisioning state.
	Provisioning Phase = "Provisioning"
	// Ready status identifies that the provider is ready to be used.
	Ready Phase = "Ready"
	// Failed status defines a failed state of provider provisioning.
	Failed Phase = "Failed"
)
