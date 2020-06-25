package main

import (
	"log"

	arieslog "github.com/hyperledger/aries-framework-go/pkg/common/log"
	"github.com/spf13/pflag"

	"github.com/scoir/canis/pkg/agent"
	"github.com/scoir/canis/pkg/controller"
	"github.com/scoir/canis/pkg/framework"
)

var agentID string

func init() {
	pflag.StringVar(&agentID, "id", "", "The unique ID for this Agent")
}

func main() {
	arieslog.SetLevel("aries-framework/out-of-band/service", arieslog.CRITICAL)
	//arieslog.SetLevel("aries-framework/did-exchange/service", arieslog.DEBUG)
	//arieslog.SetLevel("aries-framework/http", arieslog.DEBUG)
	conf, err := framework.NewFileConfig()
	if err != nil {
		log.Fatalln("error reading config", err)
	}

	ctx := conf.GetAriesContext()
	theAgent, err := agent.InitializeAgent(agentID, conf)
	if err != nil {
		log.Fatalln("error initializing agent", err)
	}

	runner, err := controller.New(
		ctx,
		conf.GRPC.Host,
		conf.GRPC.Port,
		conf.GRPCBridge.Host,
		conf.GRPCBridge.Port,
		theAgent)

	if err != nil {
		log.Fatalln("unable to start steward", err)
	}

	err = runner.Launch()
	if err != nil {
		log.Fatalln("launch errored with", err)
	}

	log.Println("Shutdown")
}
