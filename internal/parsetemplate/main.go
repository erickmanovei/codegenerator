package parsetemplate

import (
	"bytes"
	"os"
	"text/template"
)

type TemplateStruct struct {
	tmpl *template.Template
	buffer bytes.Buffer
}

func Init(templateFile string) *TemplateStruct {
	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		println("err 1")
		panic(err)
	}

	return &TemplateStruct{tmpl: tmpl}
}

func (t *TemplateStruct) Parse(data interface{}) {
	error := t.tmpl.Execute(&t.buffer, data)
	if error != nil {
		println("error 2")
		panic(error)
	}
}

func (t *TemplateStruct) ToString() string {
	return t.buffer.String()
}

func (t *TemplateStruct) Save(fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		println("error 3")
		panic(err)
	}
	defer file.Close()
	file.WriteString(t.buffer.String())
}