package cmd

import (
	"github.com/twhiston/pb/pb"
	"io"
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
