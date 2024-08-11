package help

import "fmt"

func Help() {
	helpText := `
Command Line TODO App

This is a simple CLI based Todo application using GO for managing your TODO list.

Commands

    list    - Display all TODO items
    create  - Add a new TODO item
    update  - Rename an existing TODO item
    delete  - Remove a TODO item
    search  - Find TODO items by keyword

Options

    list:
        -k | --key - Display todos with time matching the given value
    Create:
        -hd | --heading -   heading value of the new todo task
        -d | --description -   Description value of the new todo task
        -t | --time -   Deadline time of the new todo task

Usage:
  todo <command> [arguments]

Examples:
  todo list                                            # Display all TODO items
  todo list -k "24-05-2023"                            # Display all TODO items matching given key
  todo create -hd "heading -d "New Task" -t "time"     # Add a new TODO item with the name "New Task"
  todo update 1 "Updated Task"                         # Rename the TODO item with ID 1 to "Updated Task"
  todo delete 1                                        # Remove the TODO item with ID 1
  todo search "Task"                                   # Find TODO items containing the word "Task"
`
	fmt.Println(helpText)
}
