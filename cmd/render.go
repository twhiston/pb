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
	"errors"
	"github.com/spf13/cobra"
	"os"
)

// renderCmd represents the render command
var renderCmd = &cobra.Command{
	Use:   "render",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {

		confFile, err := cmd.PersistentFlags().GetString("conf")
		if err != nil {
			return err
		}
		if _, err := os.Stat(confFile); os.IsNotExist(err) {
			return errors.New("Config file could not be found")
		}

		return nil
	},
}

func init() {
	RootCmd.AddCommand(renderCmd)

	renderCmd.PersistentFlags().String("dir", "", "A directory to generate a block (or blocks from). Defaults to cwd. If set this overrides conf and c flags")
	renderCmd.PersistentFlags().String("conf", "", "A config yml to generate a block from")
	renderCmd.PersistentFlags().String("c", "", "Set the c file that the block will use. Of blank looks for a file called ")

}
