package main

import (
	// "fmt"
	"os"
	"encoding/csv"
	"err"
	c "github.com/tharun13055/golang/todo/action/create"

	h "github.com/tharun13055/golang/todo/action/help"
)

func main() {
	if len(os.Args) < 2  {
		h.Help()
		return
	}

	command_one:= os.Args[1]


	if len(os.Args) < 3 {
		switch command_one {
			case "-h","--help":
				h.Help()
				return
			case "-c","--create":
				c.Create()
				return
			default:
			  a.Help()
			  return
		}
	}

}

