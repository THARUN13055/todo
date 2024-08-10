package help

import "fmt"

func Help(){
    helpText := `
# Command Line TODO App

This is a simple CLI based Todo application using GO for managing your TODO list.

## Commands

list    - Display all TODO items
create  - Add a new TODO item
update  - Rename an existing TODO item
delete  - Remove a TODO item
search  - Find TODO items by keyword

Usage:
  todo <command> [arguments]

Examples:
  todo list                  # Display all TODO items
  todo create "New Task"     # Add a new TODO item with the name "New Task"
  todo update 1 "Updated Task" # Rename the TODO item with ID 1 to "Updated Task"
  todo delete 1              # Remove the TODO item with ID 1
  todo search "Task"         # Find TODO items containing the word "Task"

Use "todo help" for more information about a command.
`
    fmt.Println(helpText)
}