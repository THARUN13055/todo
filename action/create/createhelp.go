package createhelp

func CreateHelp(){

	createHelp := `

	Main command:
	  ./todo --create or -c <value>

	  -- create "This is first task" --discription or -d "Here we need to bref about it"
	`

	fmt.Println(createhelp)
}