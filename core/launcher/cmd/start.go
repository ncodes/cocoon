// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/ncodes/cocoon/core/launcher/ccode"
	"github.com/ncodes/cocoon/core/launcher/config"
	logging "github.com/op/go-logging"
	"github.com/spf13/cobra"
)

func init() {
	config.ConfigureLogger()
}

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the launcher",
	Long:  `Starts the launcher. Pulls cocoon code specified in COCOON_CODE_URL, builds it and runs it.`,
	Run: func(cmd *cobra.Command, args []string) {

		doneCh := make(chan bool, 1)
		ccInstallFailedCh := make(chan bool, 1)

		var log = logging.MustGetLogger("launcher")

		// install cooncode
		ccode.Install(ccInstallFailedCh)

		if <-ccInstallFailedCh {
			log.Fatal("aborting: cocoon code installation failed")
		}

		// start cocoon code
		ccode.Start()

		<-doneCh
	},
}

func init() {
	RootCmd.AddCommand(startCmd)
	// startCmd.Flags().StringP("port", "p", "5500", "The port to run bind to")
}