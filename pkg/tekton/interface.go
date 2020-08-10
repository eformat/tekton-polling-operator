package tekton

import (
	"encoding/json"

	triggersv1 "github.com/tektoncd/triggers/pkg/apis/triggers/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/bigkevmcd/tekton-polling-operator/pkg/git"
)

// Resolver implementations generate resources from the combination of template
// and binding references, within a namespace, and applying the body of a
// commit.
type Resolver interface {
	Resolve(ns string, bindings []*triggersv1.EventListenerBinding, template triggersv1.EventListenerTemplate, commit git.Commit) ([]json.RawMessage, error)
}

// ResolverClient is used by a resolver to get the necessary resources to
// generate the resources.
type ResolverClient interface {
	GetTriggerBinding(name string, options metav1.GetOptions) (*triggersv1.TriggerBinding, error)
	GetTriggerTemplate(name string, options metav1.GetOptions) (*triggersv1.TriggerTemplate, error)
	GetClusterTriggerBinding(name string, options metav1.GetOptions) (*triggersv1.ClusterTriggerBinding, error)
}