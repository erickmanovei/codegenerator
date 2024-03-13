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
}
