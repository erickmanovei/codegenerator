package model

import (
	"github.com/erickmanovei/codegenerator/internal/parseargs"
	"github.com/erickmanovei/codegenerator/internal/parsetemplate"
)

type model struct {
	Name string
	Age int
}

func Model(data *parseargs.Args) {
	tmpl := parsetemplate.Init("internal/model/files/model.tmpl")
	tmpl.Parse(&model{Name: "Robson Fernandes", Age: 50})
	tmpl.Save("internal/model/files/" + data.Arg + ".ts")
}
