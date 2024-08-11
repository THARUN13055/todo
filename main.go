package main

import (
	"os"

	c "github.com/tharun13055/golang/todo/action/create"
	h "github.com/tharun13055/golang/todo/action/help"
	l "github.com/tharun13055/golang/todo/action/list"
	Checkenv "github.com/tharun13055/golang/todo/action/checkenv"
)

func main() {
	if len(os.Args) < 2 {
		h.Help()
		return
	}

	command_one := os.Args[1]
	Checkenv.CreateFile()

	if len(os.Args) > 1 {
		switch command_one {
		case "-h", "--help":
			h.Help()
			return
		case "c", "create":
			c.Create()
			return
		case "l", "list":
			l.List()
		default:
			h.Help()
			return
		}
	}

}
