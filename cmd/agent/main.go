package main

import (
	arieslog "github.com/hyperledger/aries-framework-go/pkg/common/log"

	"github.com/scoir/canis/pkg/agent/cmd"
)

func main() {
	arieslog.SetLevel("aries-framework/out-of-band/service", arieslog.CRITICAL)
	//arieslog.SetLevel("aries-framework/did-exchange/service", arieslog.DEBUG)
	//arieslog.SetLevel("aries-framework/http", arieslog.DEBUG)
	cmd.Execute()
}
