package cmd

import (
	"bufio"
	"errors"
	"github.com/twhiston/pb/pb"
	"io"
	"log"
	"text/template"
)

func renderSchemaTemplate(wr io.Writer, config *pb.Config) error {
	//Load the template, parse and render it
	filename := config.Name + ".yml"
	tmpl := template.New(filename) //create a new template

	data, err := pb.Asset("tpl/schema.yml.tpl")
	if err != nil {
		return err
	}

	if tmpl, err = tmpl.Parse(string(data[:])); err != nil {
		return err
	}
	err = tmpl.Execute(wr, config)
	if err != nil {
		return err
	}
	return nil
}

// getInput takes an id, printed in the user prompt, and asks the user for input
// It does not allow empty inputs, and will prompt the user again if they fail to enter something valid
func getInput(id string, def string, r io.Reader) (string, error) {
	output, err := getInputImpl(id, def, r, 0, 10)
	if err != nil {
		return "", err
	}
	return output, nil
}

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
		depth += 1
		return getInputImpl(id, def, r, depth, max)
	}
	return token, nil

}
