// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	externaldnsv1 "github.com/nginxinc/kubernetes-ingress/v3/pkg/apis/externaldns/v1"
	versioned "github.com/nginxinc/kubernetes-ingress/v3/pkg/client/clientset/versioned"
	internalinterfaces "github.com/nginxinc/kubernetes-ingress/v3/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/nginxinc/kubernetes-ingress/v3/pkg/client/listers/externaldns/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// DNSEndpointInformer provides access to a shared informer and lister for
// DNSEndpoints.
type DNSEndpointInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.DNSEndpointLister
}

type dNSEndpointInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewDNSEndpointInformer constructs a new informer for DNSEndpoint type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewDNSEndpointInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredDNSEndpointInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredDNSEndpointInformer constructs a new informer for DNSEndpoint type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredDNSEndpointInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ExternaldnsV1().DNSEndpoints(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ExternaldnsV1().DNSEndpoints(namespace).Watch(context.TODO(), options)
			},
		},
		&externaldnsv1.DNSEndpoint{},
		resyncPeriod,
		indexers,
	)
}

func (f *dNSEndpointInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredDNSEndpointInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *dNSEndpointInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&externaldnsv1.DNSEndpoint{}, f.defaultInformer)
}

func (f *dNSEndpointInformer) Lister() v1.DNSEndpointLister {
	return v1.NewDNSEndpointLister(f.Informer().GetIndexer())
}
