//+build wireinject

package agent

import (
	"github.com/google/wire"

	"github.com/scoir/canis/pkg/framework"
)

func InitializeAgent(agentID string, conf *framework.Config) (*Agent, error) {
	wire.Build(
		NewAgent)
	return &Agent{}, nil
}
