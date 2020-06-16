package context

import "k8s.io/client-go/kubernetes"

type Clientset struct {
	*kubernetes.Clientset
	Namespace string
}
