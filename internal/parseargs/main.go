package parseargs

import (
	"log"
	"os"
)

type Args struct {
	Command string
	Arg     string
}

func Parse() *Args {
	if len(os.Args) < 3 {
		log.Fatal("No arguments provided")
	}
	command := os.Args[1]
	arg := os.Args[2]
	return &Args{Command: command, Arg: arg}
}
