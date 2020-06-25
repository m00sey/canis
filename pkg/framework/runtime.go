package framework

import (
	"github.com/pkg/errors"

	"github.com/scoir/canis/pkg/runtime"
)

func (r *Config) Executor() (runtime.Executor, error) {
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
	case "docker":
	}

	return r.exec, errors.Wrap(err, "unable to launch runtime from config")

}

func (r *Config) loadK8s() (runtime.Executor, error) {

	return nil, nil
}
