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
	"strconv"

	"github.com/spf13/cobra"

	vebtree "github.com/matchstick/vebtree/lib"
)

func newInsertCmd() *cobra.Command {
	// insertCmd represents the insert command.
	var insertCmd = &cobra.Command{
		Use:   "insert",
		Short: "",
		Long:  ` `,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			var vals []int
			for _, arg := range args {
				val, err := strconv.Atoi(arg)
				if err != nil {
					fmt.Printf("%v\n", err)
					return
				}
				fmt.Printf("%d\n", val)
				vals = append(vals, val)
			}
			vebtree.Insert(vals)
		},
	}

	return insertCmd
}
