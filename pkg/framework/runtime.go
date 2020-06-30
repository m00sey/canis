package framework

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/pkg/errors"

	"github.com/scoir/canis/pkg/runtime"
	"github.com/scoir/canis/pkg/runtime/kubernetes"
	"github.com/scoir/canis/pkg/runtime/proc"
)

type RuntimeConfig struct {
	Runtime    string             `mapstructure:"runtime"`
	Kubernetes *kubernetes.Config `mapstructure:"kubernetes"`
	Proc       *proc.Config       `mapstructure:"proc"`

	lock sync.Mutex
	exec runtime.Executor
}

func (r *RuntimeConfig) Executor() (runtime.Executor, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	if r.exec != nil {
		return r.exec, nil
	}

	var err error
	switch r.Runtime {
	case "kubernetes":
		r.exec, err = r.loadK8s()
	case "proc":
		r.exec, err = r.loadProc()
	case "docker":
	}

	return r.exec, errors.Wrap(err, "unable to launch runtime from config")

}

func (r *RuntimeConfig) loadK8s() (runtime.Executor, error) {

	return nil, nil
}

func (r *RuntimeConfig) loadProc() (runtime.Executor, error) {
	d, _ := json.MarshalIndent(r.Proc, " ", " ")
	fmt.Println(string(d))
	return proc.New(r.Proc), nil
}
