/*
Copyright © 2024 Steven A. Zaluk

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
	"github.com/spf13/cobra"
	"github.com/stevezaluk/go-protocol-poc/src"
)

var server src.Server

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start server side routine",
	Long:  `Executing the 'server' command will start the server and will wait for incoming client connections`,
	Run: func(cmd *cobra.Command, args []string) {
		server.Start()
		server.AcceptConnections()
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		server.Stop()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
