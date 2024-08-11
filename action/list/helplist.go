package list

import "fmt"

func CreateHelp() {

	createHelp := `
	Command: 
	  todo --list or -l  # It will list all the tasks

	  todo --list or -l  --key or -k # It will list specific task

	Args:
	  --list or -l # Here the l is meaning for list which list all the data which you posted
	  -- key or -k # Here the k is meaning for key. For example in time (--time) you mension as tommrow 
	      command:  todo -l -k tommow 
		  output:
		    4 "biscut" "buy the borbom" "tommrow" 
			6 "sunset" "going for beach" "tommrow"
	`

	fmt.Println(createHelp)
}
