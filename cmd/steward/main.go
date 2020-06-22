package main

import (
	"log"

	arieslog "github.com/hyperledger/aries-framework-go/pkg/common/log"

	"github.com/scoir/canis/pkg/controller"
	"github.com/scoir/canis/pkg/framework"
	"github.com/scoir/canis/pkg/steward"
)

func main() {
	arieslog.SetLevel("aries-framework/out-of-band/service", arieslog.CRITICAL)
	//arieslog.SetLevel("aries-framework/did-exchange/service", arieslog.DEBUG)
	//arieslog.SetLevel("aries-framework/http", arieslog.DEBUG)
	conf, err := framework.NewFileConfig()
	if err != nil {
		log.Fatalln("error reading config", err)
	}

	ctx := conf.GetAriesContext()

	agent, err := steward.InitializeAgent(ctx, conf)
	if err != nil {
		log.Fatalln("error initializing agent", err)
	}

	runner, err := controller.New(
		ctx,
		conf.GRPC.Host,
		conf.GRPC.Port,
		conf.GRPCBridge.Host,
		conf.GRPCBridge.Port,
		agent)

	if err != nil {
		log.Fatalln("unable to start controller", err)
	}

	err = runner.Launch()
	if err != nil {
		log.Fatalln("launch errored with", err)
	}

	log.Println("Shutdown")
}
