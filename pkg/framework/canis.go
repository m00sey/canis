package framework

import (
	"log"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/scoir/canis/pkg/framework/context"
)

func GetClientSet(kubeconfig, namespace string) *context.Clientset {
	var config *rest.Config
	var err error

	if kubeconfig == "" {
		config, err = rest.InClusterConfig()
		if err != nil {
			log.Fatalln("error building in cluster config", err)
		}
	} else {
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			log.Fatalln("error building local config from file", err)
		}
	}

	cs, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	return &context.Clientset{Clientset: cs, Namespace: namespace}
}

func GetClientSetWithConfig(config *Config) *context.Clientset {
	return GetClientSet(config.KubeConfig, config.Namespace)
}
