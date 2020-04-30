// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	time "time"

	hivev1 "github.com/openshift/hive/pkg/apis/hive/v1"
	versioned "github.com/openshift/hive/pkg/client/clientset/versioned"
	internalinterfaces "github.com/openshift/hive/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/openshift/hive/pkg/client/listers/hive/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ClusterStateInformer provides access to a shared informer and lister for
// ClusterStates.
type ClusterStateInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ClusterStateLister
}

type clusterStateInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewClusterStateInformer constructs a new informer for ClusterState type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewClusterStateInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredClusterStateInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredClusterStateInformer constructs a new informer for ClusterState type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredClusterStateInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.HiveV1().ClusterStates(namespace).List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.HiveV1().ClusterStates(namespace).Watch(options)
			},
		},
		&hivev1.ClusterState{},
		resyncPeriod,
		indexers,
	)
}

func (f *clusterStateInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredClusterStateInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *clusterStateInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&hivev1.ClusterState{}, f.defaultInformer)
}

func (f *clusterStateInformer) Lister() v1.ClusterStateLister {
	return v1.NewClusterStateLister(f.Informer().GetIndexer())
}