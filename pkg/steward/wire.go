//+build wireinject

package steward

import (
	"github.com/google/wire"
	"github.com/hyperledger/aries-framework-go/pkg/framework/aries/api"

	"github.com/scoir/canis/pkg/framework"
	"github.com/scoir/canis/pkg/schema"
)

func InitializeAgent(ctx api.Provider, conf *framework.Config) (*Steward, error) {
	wire.Build(
		New,
		schema.New)
	return &Steward{}, nil
}
