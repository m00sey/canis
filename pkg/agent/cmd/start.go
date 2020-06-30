/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/scoir/canis/pkg/agent"
	"github.com/scoir/canis/pkg/controller"
)

var agentID *string

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the agent",
	Long:  `Starts the agent`,
	Run:   runStart,
}

func runStart(_ *cobra.Command, _ []string) {

	if *agentID == "" {
	}
	ctx := config.GetAriesContext()

	a, err := agent.NewAgent(*agentID, config)
	if err != nil {
		log.Fatalln("error initializing agent", err)
	}

	runner, err := controller.New(
		ctx,
		config.GRPC.Host,
		config.GRPC.Port,
		config.GRPCBridge.Host,
		config.GRPCBridge.Port,
		a)

	if err != nil {
		log.Fatalln("unable to start steward", err)
	}

	err = runner.Launch()
	if err != nil {
		log.Fatalln("launch errored with", err)
	}

	log.Println("Shutdown")
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().StringVar(agentID, "id", "", "The unique ID of this agent")
	_ = startCmd.MarkFlagRequired("agentID")
}
