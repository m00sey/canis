package steward

import (
	"github.com/hyperledger/aries-framework-go/pkg/didcomm/messaging/msghandler"
	"github.com/hyperledger/aries-framework-go/pkg/didcomm/transport/ws"
	"github.com/hyperledger/aries-framework-go/pkg/framework/aries"
	"github.com/hyperledger/aries-framework-go/pkg/framework/aries/defaults"
	"github.com/hyperledger/aries-framework-go/pkg/framework/context"
	"github.com/hyperledger/aries-framework-go/pkg/storage/leveldb"

	"github.com/scoir/canis/pkg/framework"
)

type Config struct {
	ctx                   *context.Provider `yaml:"-"`
	framework.AgentConfig `yaml:",inline"`
	Datastore             framework.DatastoreConfig `yaml:"datastore"`
	Execution             framework.RuntimeConfig   `yaml:"execution"`
}

func NewConfig() *Config {
	c := &Config{}

	c.AgentConfig.GetAriesOptions = c.getOptions
	return c
}

func (r *Config) getOptions() []aries.Option {
	wsinbound := r.WSInbound.Address()

	return []aries.Option{
		aries.WithMessageServiceProvider(msghandler.NewRegistrar()),
		aries.WithStoreProvider(leveldb.NewProvider(r.DBPath)),
		defaults.WithInboundWSAddr(wsinbound, wsinbound),
		aries.WithOutboundTransports(ws.NewOutbound()),
		aries.WithServiceEndpoint(r.Address()),
	}
}
