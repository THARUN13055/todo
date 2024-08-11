package main

import (
	"os"

	c "github.com/tharun13055/golang/todo/action/create"
	h "github.com/tharun13055/golang/todo/action/help"
	Checkenv "github.com/tharun13055/todo/action/checkenv"
)

func main() {
	if len(os.Args) < 2 {
		h.Help()
		return
	}

	command_one := os.Args[1]
	Checkenv.CreateFile()

	if len(os.Args) < 3 {
		switch command_one {
		case "-h", "--help":
			h.Help()
			return
		case "-c", "--create":
			c.Create()
			return
		default:
			h.Help()
			return
		}
	}

}
