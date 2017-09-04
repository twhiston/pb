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
	"bytes"
	"errors"
	"github.com/spf13/cobra"
	"github.com/twhiston/pb/pb"
	"io/ioutil"
	"os/user"
	"time"
)

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "new",
	Short: "Generates a patchblock project skeleton",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Bare repos get a simple build
		isBare, err := cmd.PersistentFlags().GetBool("bare")
		if err != nil {
			return err
		}
		if isBare == true {
			return handleBareBuild(args)
		}

		//If its not a bare build we go into 'interactive console' mode and ask a load of questions

		return nil
	},
}

func handleBareBuild(args []string) error {
	if len(args) == 0 {
		return errors.New("Must pass in patchblock name as argument when creating a proect skeleton")
	}

	config, err := makeBareConfigData(args)
	if err != nil {
		return err
	}

	//Render the template to a buffer and write it to a file
	buf := bytes.NewBuffer(*new([]byte))
	renderSchemaTemplate(buf, config)
	return ioutil.WriteFile(config.Name+".yml", buf.Bytes(), 0644)
}

func makeBareConfigData(args []string) (*pb.Config, error) {
	//Make a config object
	config := new(pb.Config)

	config.Name = string(args[0])

	//Current User
	usr, err := user.Current()
	if err != nil {
		return nil, err
	}
	config.Username = usr.Username

	//Time
	t := time.Now()
	config.DateNow = t.Format("Mon 22 Jan 2006 15:04:05")

	//We actually want to add 1 element for input, output and var for example, but we just add blank strings
	config.Inputs = make([]pb.Input, 1)
	config.Outputs = make([]pb.Output, 1)
	config.Vars = make([]pb.Vars, 1)

	config.Code.Block = "	# enclose a block of code in backticks for multi line compatibility"
	config.Code.Path = "	# full path to the .c file you wish to parse into the block definition on render. If path is blank and no block tries cwd"

	return config, nil
}

func init() {
	RootCmd.AddCommand(genCmd)

	genCmd.PersistentFlags().Bool("bare", false, "If true will only generate skeletons with no questions or data etc.... MUST have patchblock name as argument")

}
