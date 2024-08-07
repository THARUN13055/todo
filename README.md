# Command Line TODO App

This is a simple CLI based Todo application using GO for managing your TODO list.

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

+ List all TODO items:
   
   `./todo list`
   

+ Add a new TODO item:
   
   `./todo new "Buy groceries"`
   

+ Rename a TODO item:
   
   `./todo saveas "Buy groceries" "Purchase food items"`
   

+ Delete a TODO item:
   
   `./todo delete "Purchase food items"`
   

+ Search for TODO items:
   
   `./todo search "food"`
   

## Getting Started

+ Clone this repository
+ Make sure you have the necessary permissions to execute the script
+ Run the app using the syntax described above

Happy organizing!
