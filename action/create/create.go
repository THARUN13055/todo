package create

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	createfile "github.com/tharun13055/golang/todo/action/create"
)

type storescreate struct {
	Header      string
	Description string
	Duration    string
}

func Create() {
	// Create the file and Header
	createfile.CreateFile()

	// Check if the necessary arguments are provided
	if len(os.Args) <= 2 {
		fmt.Println("Create your own text :)")
		createfile.createHelp()
		os.Exit(1)
	}

	// Initialize the struct
	var storecreate storescreate
	storecreate.Header = os.Args[3]
	storecreate.Description = os.Args[5]
	storecreate.Duration = os.Args[7]

	// Check if Header (Topic) is provided
	if storecreate.Header == "" {
		fmt.Println("Topic is Mandatory :)")
		createfile.createHelp()
		os.Exit(1)
	}

	// Check and set Description
	if os.Args[4] == "-d" || os.Args[4] == "--description" {
		if storecreate.Description == "" {
			storecreate.Description = "Null"
		}
	}

	// Check and set Duration
	if os.Args[6] == "-t" || os.Args[6] == "--time" {
		if storecreate.Duration == "" {
			storecreate.Duration = "Null"
		}
	}

	index := generateindex()

	createonetask := []string{fmt.Sprint(index), storecreate.Header, storecreate.Description, storecreate.Duration}

	// You can now use createonetask to write to the CSV file or perform other operations
}

// Function to generate index
func generateindex() int {
	// Open the CSV file
	file, err := os.Open("data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)

	// Read all rows
	rows, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Calculate the new index
	numRows := len(rows) - 1
	return numRows + 1
}
