package pb

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

type Variable struct {
	Info  string `yaml:"info"`
	Value string `yaml:"value"`
}

type Input struct {
	Variable
	Editable bool `yaml:"editable"`
}

type Output struct {
	Variable
}

type Vars struct {
	Dtype string `yaml:"dtype"`
	Name  string `yaml:"name"`
}

type Code struct {
	Block string `yaml:"block"`
	Path  string `yaml:"path"`
}

type Config struct {
	Name     string   `yaml:"name"`
	Category string   `yaml:"category"`
	Info     string   `yaml:"info"`
	Inputs   []Input  `yaml:"inputs"`
	Outputs  []Output `yaml:"outputs"`
	Vars     []Vars   `yaml:"vars"`
	Code     Code     `yaml:"code"`
	Help     string   `yaml:"help"`
	InternalTemplateData
}

type InternalTemplateData struct {
	Username string
	DateNow  string
}
