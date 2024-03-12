package main

import (
	"log"

	"github.com/erickmanovei/codegenerator/internal/model"
	"github.com/erickmanovei/codegenerator/internal/parseargs"
)

func main() {
	data := parseargs.Parse()
	switch data.Command {
	case "model":
		model.Model(data)
	default:
		log.Fatal("Unknown command")
	}
	// var file string
	// flag.StringVar(&file, "file", "default", "file to read")
	// flag.Parse()
	// println(os.Args[1])
	// println(file)
}
