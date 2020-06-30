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

	"github.com/scoir/canis/pkg/controller"
	"github.com/scoir/canis/pkg/schema"
	"github.com/scoir/canis/pkg/steward"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the steward orchestration service",
	Long:  `Starts a steward orchestration service`,
	Run:   runStart,
}

func runStart(cmd *cobra.Command, args []string) {
	ctx := config.GetAriesContext()

	client := schema.New()
	agent, err := steward.New(ctx, config, client)
	if err != nil {
		log.Fatalln("error initializing steward agent", err)
	}

	runner, err := controller.New(
		ctx,
		config.GRPC.Host,
		config.GRPC.Port,
		config.GRPCBridge.Host,
		config.GRPCBridge.Port,
		agent)

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
}
