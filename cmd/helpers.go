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
	"bufio"
	"errors"
	"github.com/twhiston/pb/pb"
	"io"
	"log"
	"strings"
	"text/template"
)

//Render a .yml block template from a pb.Config struct
func renderSchemaTemplate(wr io.Writer, config *pb.Config) error {
	//Load the template, parse and render it
	filename := config.Name + ".yml"
	tmpl := template.New(filename) //create a new template

	data, err := pb.Asset("assets/tpl/schema.yml.tpl")
	if err != nil {
		return err
	}

	if tmpl, err = tmpl.Parse(string(data[:])); err != nil {
		return err
	}
	return tmpl.Execute(wr, config)
}

//Render a .c file template from a pb.Config struct
func renderCTemplate(wr io.Writer, config *pb.Config) error {
	//Load the template, parse and render it
	filename := config.Name + ".c"
	tmpl := template.New(filename) //create a new template

	data, err := pb.Asset("assets/tpl/block.c.tpl")
	if err != nil {
		return err
	}

	if tmpl, err = tmpl.Parse(string(data[:])); err != nil {
		return err
	}
	return tmpl.Execute(wr, config)
}

func renderXMLTemplate(wr io.Writer, config *pb.Config) error {
	//Load the template, parse and render it
	filename := config.Name + ".xml"
	tmpl := template.New(filename) //create a new template

	data, err := pb.Asset("assets/tpl/block.xml.tpl")
	if err != nil {
		return err
	}

	if tmpl, err = tmpl.Parse(string(data[:])); err != nil {
		return err
	}
	return tmpl.Execute(wr, config)
}

func SanitiseStringInput(i string) string {
	return strings.Replace(i, " ", "_", -1)
}

// getInput takes an id, printed in the user prompt, and asks the user for input
// It does not allow empty inputs, and will prompt the user again if they fail to enter something valid
func getInput(id string, def string, r io.Reader) (string, error) {
	return getInputImpl(id, def, r, 0, 10)
}

// getInputImpl should not be used directly, call getInput instead
func getInputImpl(id string, def string, r io.Reader, depth int, max int) (string, error) {

	if depth >= max {
		return "", errors.New("Too many attempts")
	}
	if def != "" {
		log.Print(id, ": [", def, "]")
	} else {
		log.Print(id, ":")
	}

	scanner := bufio.NewScanner(r)
	scanner.Scan()
	token := scanner.Text()
	if def != "" && token == "" {
		token = def
	} else if token == "" {
		log.Println(id, "cannot be blank, do it again and this time don't embarrass yourself")
		depth++
		return getInputImpl(id, def, r, depth, max)
	}
	return token, nil

}
