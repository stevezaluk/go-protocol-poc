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

var client src.Client

var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Start client side routine",
	Long: `Executing this with the 'client' command will initiate a connection to the server 
and provide an interface for sending messages. Executing this with no arguments will send a basic message.
Any additional arguments passed will be processed as messages and will be sent to the server`,
	Run: func(cmd *cobra.Command, args []string) {
		uri, _ := cmd.Flags().GetString("uri")

		client.Uri = uri

		client.Connect()
		client.NegotiateKeys()
		client.ValidatePublicKey()
		client.SendWelcome()
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		// client.Disconnect()
	},
}

func init() {
	rootCmd.AddCommand(clientCmd)
	clientCmd.PersistentFlags().StringP("uri", "u", "localhost:8000", "The URI you want to try and connect to")
}
