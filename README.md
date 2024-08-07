# Command Line TODO App

This is a simple command-line application for managing your TODO list.

## Commands

| Command | Description |
|---------|-------------|
| list    | Display all TODO items |
| create  | Add a new TODO item |
| update  | Rename an existing TODO item |
| delete  | Remove a TODO item |
| search  | Find TODO items by keyword |

## Usage

The general syntax for using the app is:


./todo <command> [arguments]


### Examples

1. List all TODO items:
   
   ./todo list
   

2. Add a new TODO item:
   
   ./todo new "Buy groceries"
   

3. Rename a TODO item:
   
   ./todo saveas "Buy groceries" "Purchase food items"
   

4. Delete a TODO item:
   
   ./todo delete "Purchase food items"
   

5. Search for TODO items:
   
   ./todo search "food"
   

## Getting Started

1. Clone this repository
2. Make sure you have the necessary permissions to execute the script
3. Run the app using the syntax described above

Happy organizing!