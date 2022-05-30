/*
Copyright © 2020 Michael Rubin <mhr@neverthere.org>

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
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func newVersionCmd() *cobra.Command {
	const numVersionCmdArgs = 0

	return &cobra.Command{
		Use:   "version",
		Short: "Output release version.",
		Long:  `Outputs release version. `,
		Args:  cobra.MaximumNArgs(numVersionCmdArgs),
		Run: func(cmd *cobra.Command, args []string) {
			version := "v0.0.1"
			fmt.Printf("%s\n", version)
		},
	}
}

func newRootCmd() *cobra.Command {
	// rootCmd represents the base command when called without any subcommands.
	return &cobra.Command{
		Use:   "vebtree",
		Short: "Datastrucutre cmd",
		Long: `Fun vebtree. `,
	}
}

const exitErr = 1

// Execute is the single function to set up the complete set of nested cobra
// based commands that provide CLI functionality to the exifsort library.
func Execute() {
	cobra.OnInitialize(initConfig)

	rootCmd := newRootCmd()

	rootCmd.AddCommand(newInsertCmd())
	rootCmd.AddCommand(newDeleteCmd())
	rootCmd.AddCommand(newLookupCmd())
	rootCmd.AddCommand(newNextCmd())
	rootCmd.AddCommand(newPrevCmd())
	rootCmd.AddCommand(newVersionCmd())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(exitErr)
	}
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
}
