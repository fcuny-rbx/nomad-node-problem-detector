/*
Copyright 2021 Roblox Corporation

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

package main

import (
	log "github.com/sirupsen/logrus"
	"os"
	"time"

	aggregator "github.com/nomad-node-problem-detector/aggregator"
	config "github.com/nomad-node-problem-detector/config"
	detector "github.com/nomad-node-problem-detector/detector"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:                 "npd",
		Usage:                "Nomad node problem detector",
		Version:              "v1.0.0",
		Compiled:             time.Now(),
		EnableBashCompletion: true,
		Authors: []*cli.Author{
			{
				Name:  "Shishir Mahajan",
				Email: "smahajan@roblox.com",
			},
		},
		Commands: []*cli.Command{
			aggregator.AggregatorCommand,
			detector.DetectorCommand,
			config.ConfigCommand,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
