package tekton

import (
	"context"

	triggersv1 "github.com/tektoncd/triggers/pkg/apis/triggers/v1alpha1"
	triggersclientset "github.com/tektoncd/triggers/pkg/client/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// VersionedClient is an implementation of the ResolverClient that uses a Tekton
// Triggers versioned client to get the resources.
type VersionedClient struct {
	ns     string
	client triggersclientset.Interface
}

// NewVersionedClient creates and returns a new VersionedClient.
func NewVersionedClient(ns string, c triggersclientset.Interface) *VersionedClient {
	return &VersionedClient{
		ns:     ns,
		client: c,
	}
}

// GetTriggerBinding is an implementation of the ResolverClient interface.
func (v *VersionedClient) GetTriggerBinding(name string) (*triggersv1.TriggerBinding, error) {
	return v.client.TriggersV1alpha1().TriggerBindings(v.ns).Get(context.Background(), name, metav1.GetOptions{})
}

// GetTriggerTemplate is an implementation of the ResolverClient interface.
func (v *VersionedClient) GetTriggerTemplate(name string) (*triggersv1.TriggerTemplate, error) {
	return v.client.TriggersV1alpha1().TriggerTemplates(v.ns).Get(context.Background(), name, metav1.GetOptions{})
}

// GetTriggerTriggerBinding is an implementation of the ResolverClient interface.
func (v *VersionedClient) GetClusterTriggerBinding(name string) (*triggersv1.ClusterTriggerBinding, error) {
	return v.client.TriggersV1alpha1().ClusterTriggerBindings().Get(context.Background(), name, metav1.GetOptions{})
}
