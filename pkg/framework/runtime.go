package framework

import (
	"github.com/pkg/errors"

	"github.com/scoir/canis/pkg/runtime"
)

func (r *Config) Executor() (runtime.Executor, error) {
	return nil, errors.New("implement me")
}
