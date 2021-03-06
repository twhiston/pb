package pb

import "encoding/xml"

///////////////////////////////////
// CONFIG
///////////////////////////////////
// 	name:
// 	category:
// 	info:
// 	inputs:
//	 	- info:
//     	  value:      #optional
//     	  editable:   #optional
// 	outputs:
// 		- info:
//     	  value:      #optional
// 	vars:
//   	- dtype:
// 		  name:
//	code:			  #Either must be set
//		block:
//		path:
///////////////////////////////////

// Variable structs form the base fields of all items in the patchblock data section
type Variable struct {
	Info  string `yaml:"info"`
	Value string `yaml:"value"`
}

// Input describes a patchblock input port
type Input struct {
	Variable `yaml:",inline"`
	Editable string `yaml:"editable"`
}

//Output describles a patchblock output port
type Output struct {
	Variable `yaml:",inline"`
}

// Var describes a patchblock variable, Dtype must be a valid C type name. Name should include array size if type is array
// as on rendering the variable will be composed as `Dtype Name;`
type Var struct {
	Variable `yaml:",inline"`
	Dtype    string `yaml:"dtype"`
	Name     string `yaml:"name"`
}

// Code represents the location of the function CDATA code. It is suggested to set the path and use the generated C file
// however is also possible to use block and paste the code directly in the yaml, backticks should be used to surround the block
// note that adding code here will not allow you to do render time macro expansion or use the testing framework
type Code struct {
	Block string
	Path  string `yaml:"path"`
}

// Config represents the full config structure of the patchblock, which will be passed to a template for rendering,
// or will be ingested by the xml renderer when composing the final output
type Config struct {
	Name     string   `yaml:"name"`
	Category string   `yaml:"category"`
	Info     string   `yaml:"info"`
	Inputs   []Input  `yaml:"inputs"`
	Outputs  []Output `yaml:"outputs"`
	Vars     []Var    `yaml:"vars"`
	Code     Code     `yaml:"code"`
	Help     string   `yaml:"help"`
	InternalTemplateData
}

// InternalTemplateData represents any data that is not strictly needed for the patchblock, but which we need at render
// time and cannot be composed at the template level
type InternalTemplateData struct {
	Username       string
	DateNow        string
	BlockNameLower string
	ParsedFunction []string
}

type XMLConfig struct {
	XMLName  xml.Name    `xml:"block"`
	Name     string      `xml:"name,attr"`
	Category string      `xml:"category"`
	Info     string      `xml:"info"`
	Data     XMLData     `xml:"data"`
	Function XMLFunction `xml:"function"`
	Help     XMLHelp     `xml:"help"`
}

type XMLData struct {
	XMLName   xml.Name          `xml:"data"`
	Name      string            `xml:"name,attr"`
	Variables []XMLDataVariable `xml:"variable"`
}

type XMLDataVariable struct {
	XMLName  xml.Name `xml:"variable"`
	Socket   string   `xml:"socket,attr"`
	Info     string   `xml:"info,attr"`
	Editable string   `xml:"editable,attr"`
	Name     string   `xml:"name,attr"`
	Value    string   `xml:"value,attr"`
	Dtype    string   `xml:"dtype,attr"`
}

type XMLFunction struct {
	XMLName  xml.Name `xml:"function"`
	Name     string   `xml:"name,attr"`
	Function string   `xml:",cdata"`
}

type XMLHelp struct {
	XMLName xml.Name `xml:"help"`
	Help    string   `xml:",cdata"`
}
