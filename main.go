package main

import (
	"fmt"
	"os"
	a "github.com/tharun13055/golang/todo/action/help"
)



func main(){
	if len(os.Args) < 2 {
		a.Help()
	}
	
	command := os.Args[1]
	lenArgs := len(os.Args)

	switch command{
	case "help":
		if lenArgs < 3 {
			a.Help()
		}
	case "create":
		fmt.Print("create")
	case "update":
		fmt.Print("update")
	case "delete":
		fmt.Print("update")
	case "search":
		fmt.Print("update")
	case "list":
		fmt.Print("update")
	}
}