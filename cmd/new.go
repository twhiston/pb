// Copyright © 2017 Tom Whiston <tom.whiston@gmail.com>
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
	"log"
	"os"
	"os/user"
	"time"
)

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "new",
	Short: "Generates a patchblock project skeleton",
	Long: `The new command will generate a new patchblock project skeleton.
This consists of:
 patchblock_name.yml 	- A yml file comprised of all the field data and information about your patchblock.
 					  	  Blocks of text should be surrounded by backticks
 patchblock_name.c		- A valid C file where you can write your patch code.
 						  This differst slighly from a standard patchblock function CDATA as it contains the function signature.
 						  This is so that later we can add a testing framework`,
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
		return handleInteractiveBuild()
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

	return doRenderFromConfig(config)
}

func makeBareConfigData(args []string) (*pb.Config, error) {
	//Make a config object
	config := new(pb.Config)

	config.Name = string(args[0])

	//Current User
	err := addHeaderConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	//We actually want to add 1 element for input, output and var for example, but we just add blank strings
	config.Inputs = make([]pb.Input, 1)
	config.Outputs = make([]pb.Output, 1)
	config.Vars = make([]pb.Var, 1)

	config.Code.Block = "	# enclose a block of code in backticks for multi line compatibility"
	config.Code.Path = "	# full path to the .c file you wish to parse into the block definition on render. If path is blank and no block tries cwd"

	return config, nil
}

func renderSchemaToCwd(config *pb.Config) error {
	buf := bytes.NewBuffer(*new([]byte))
	err := renderSchemaTemplate(buf, config)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(config.Name+".yml", buf.Bytes(), 0644)
}

func renderCToConfigCodePath(config *pb.Config) error {
	buf := bytes.NewBuffer(*new([]byte))
	err := renderCTemplate(buf, config)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(config.Code.Path, buf.Bytes(), 0644)
}

// question is a helper struct to pack a name, output reference and default value together
type question struct {
	Question string
	Storage  *string
	Default  string
}

func handleInteractiveBuild() error {

	config := new(pb.Config)

	err := addHeaderConfig(config)
	if err != nil {
		return err
	}

	qs := []question{
		{"Name", &config.Name, ""},
		{"Category", &config.Category, "Effect"},
		{"Info", &config.Info, "Editor tooltip text"},
		{"Help", &config.Help, "Shown in the help window of the editor"},
	}

	for _, v := range qs {
		err = getSimpleQuestion(&v)
		if err != nil {
			return err
		}
	}
	//Sanitize name input
	config.Name = sanitizeStringInput(config.Name)

	// LOOP FOR INPUTS
	for {
		if len(config.Inputs) >= 4 {
			log.Println("Cannot have more than 4 input")
			break
		}
		ai, err := getInput("Add Input (y/N)", "N", os.Stdin)
		if ai != "y" && ai != "Y" {
			if len(config.Inputs) > 0 {
				break
			}
			log.Println("Must have at least 1 input")
		}
		//Input - info value editable
		inp := new(pb.Input)
		iqs := []question{
			{"Info", &inp.Info, "Tooltip Text"},
			{"Value", &inp.Value, "0"},
		}

		for _, v := range iqs {
			err = getSimpleQuestion(&v)
			if err != nil {
				return err
			}
		}

		for {
			eq := question{"Editable (y/N)", &inp.Editable, "n"}
			getSimpleQuestion(&eq)
			// This is a bit :(
			if inp.Editable == "y" || inp.Editable == "Y" {
				inp.Editable = "1"
				break
			} else if inp.Editable == "n" || inp.Editable == "N" {
				inp.Editable = "0"
				break
			}
		}

		config.Inputs = append(config.Inputs, *inp)
	}

	// LOOP FOR OUTPUTS
	for {
		if len(config.Outputs) >= 4 {
			log.Println("Cannot have more than 4 outputs")
			break
		}
		ai, err := getInput("Add Output (y/N)", "N", os.Stdin)
		if ai != "y" && ai != "Y" {
			if len(config.Outputs) > 0 {
				break
			}
			log.Println("Must have at least 1 output")
		}
		//Input - info value editable
		out := new(pb.Output)
		iqs := []question{
			{"Info", &out.Info, "Tooltip Text"},
			//TODO - Does output actually need an initial value? does it work? does anyone care?
			//{"Value", &out.Value, "0"},
		}

		out.Value = "0"
		for _, v := range iqs {
			err = getSimpleQuestion(&v)
			if err != nil {
				return err
			}
		}
		config.Outputs = append(config.Outputs, *out)
	}

	// LOOP FOR VARS
	for {
		ai, err := getInput("Add Var (y/N)", "N", os.Stdin)
		if ai != "y" && ai != "Y" {
			break
		}
		//Input - info value editable
		varg := new(pb.Var)
		iqs := []question{
			{"Info", &varg.Info, "Tooltip Text"},
			{"Value", &varg.Value, "0"},
			{"dtype", &varg.Dtype, "int32_t"},
			{"name", &varg.Name, "buffer[23]"},
		}

		for _, v := range iqs {
			err = getSimpleQuestion(&v)
			if err != nil {
				return err
			}
		}
		config.Vars = append(config.Vars, *varg)
	}

	config.Code = pb.Code{Path: "./" + config.Name + ".c"}

	return doRenderFromConfig(config)
}

func doRenderFromConfig(c *pb.Config) error {

	log.Println("Rendering C File")
	err := renderCToConfigCodePath(c)
	if err != nil {
		return err
	}
	log.Println("Rendering Schema File")
	return renderSchemaToCwd(c)

}

func getSimpleQuestion(q *question) error {

	str, err := getInput(q.Question, q.Default, os.Stdin)
	if err != nil {
		return err
	}
	*q.Storage = str
	return nil
}

func addHeaderConfig(config *pb.Config) error {
	//Current User
	usr, err := user.Current()
	if err != nil {
		return err
	}
	config.Username = usr.Username

	//Time
	t := time.Now()
	config.DateNow = t.Format("Mon 22 Jan 2006 15:04:05")
	return nil
}

func init() {
	RootCmd.AddCommand(genCmd)

	genCmd.PersistentFlags().Bool("bare", false, "If true will only generate skeletons with no questions or data etc.... MUST have patchblock name as argument")

}
