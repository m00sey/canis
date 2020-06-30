package kubernetes

import (
	"io"
	"log"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/scoir/canis/pkg/datastore"
	"github.com/scoir/canis/pkg/framework/context"
	"github.com/scoir/canis/pkg/runtime"
)

type Config struct {
	KubeConfig    string `yaml:"kubeConfig"`
	Namespace     string `yaml:"namespace"`
	FQDN          string `yaml:"FQDN"`
	ImageRegistry string `yaml:"imageRegistry"`
}

type Executor struct {
}

func New(conf *Config) runtime.Executor {
	return &Executor{}
}

func (e Executor) LaunchSteward(conf []byte) (string, error) {
	panic("implement me")
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

func (e Executor) Watch(pID string) (runtime.Watcher, error) {
	panic("implement me")
}

func (e Executor) StreamLogs(pID string) (io.ReadCloser, error) {
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

func GetClientSetWithConfig(c *Config) *context.Clientset {
	return GetClientSet(c.KubeConfig, c.Namespace)
}

func (e Executor) PS() {
	panic("implement me")
}

func (e Executor) Describe() {
	panic("implement me")
}
