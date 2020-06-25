package kubernetes

import (
	"io"
	"log"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/scoir/canis/pkg/datastore"
	"github.com/scoir/canis/pkg/framework"
	"github.com/scoir/canis/pkg/framework/context"
	"github.com/scoir/canis/pkg/runtime"
)

type Executor struct {
}

func (e Executor) LaunchAgent(agent *datastore.Agent) (string, error) {
	panic("implement me")
}

func (e Executor) AgentStatus(pID string) (runtime.Process, error) {
	panic("implement me")
}

func (e Executor) ShutdownAgent(pID string) error {
	panic("implement me")
}

func (e Executor) WatchAgent(pID string) (runtime.Watcher, error) {
	panic("implement me")
}

func (e Executor) StreamAgentLogs(pID string) (io.ReadCloser, error) {
	panic("implement me")
}

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

func GetClientSetWithConfig(config *framework.Kubernetes) *context.Clientset {
	return GetClientSet(config.KubeConfig, config.Namespace)
}
