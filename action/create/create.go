package create

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type storescreate struct {
	Header      string
	Description string
	Duration    string
}

func Create() {
	// Check if the necessary arguments are provided
	if len(os.Args) <= 2 {
		fmt.Println("Create your own text :)")
		CreateHelp()
		os.Exit(1)
	}

	// Initialize the struct
	var storecreate storescreate

	//Parse the command line args
	fmt.Println("Getting data to store the todo...")

	for i := 1; i < len(os.Args); i++ {
		switch os.Args[i] {
		case "--create", "-c":
			if i+1 < len(os.Args) {
				storecreate.Header = os.Args[i+1]
				i++
			}
		case "-d", "--description":
			if i+1 < len(os.Args) {
				storecreate.Description = os.Args[i+1]
				i++
			}
		case "-t", "--time":
			if i+1 < len(os.Args) {
				storecreate.Duration = os.Args[i+1]
				i++
			}

		}

	}

	// Validate mandatory fields
	if storecreate.Header == "" {
		fmt.Println("Topic is Mandatory :)")
		CreateHelp()
		os.Exit(1)
	}

	// Set default values if fields are empty
	if storecreate.Description == "" {
		storecreate.Description = "Null"
	}
	if storecreate.Duration == "" {
		storecreate.Duration = "Null"
	}

	index := generateindex()

	createonetask := []string{fmt.Sprint(index), storecreate.Header, storecreate.Description, storecreate.Duration}	//[1, "fsdf", "fsdfs", "fsdfs"]
	fmt.Printf("Storing data: %v\n", createonetask)

	file, err := os.OpenFile("data.csv", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	log.Println("Opened data.csv file... ",err)
	if err != nil {
		log.Fatal("Falied to open data.csv",err)
	}

	defer file.Close()

	// Create a new CSV writer
	writer := csv.NewWriter(file)
	reader := csv.NewReader(file)

	//Reading data.csv
	log.Println("Reading from file, ", err)
	recodrs, err := reader.ReadAll()
	fmt.Println(recodrs)
	err = writer.Write(createonetask)
	log.Println("Writing to file, ", err)
	if err != nil {
		log.Fatal("Falied to write: ", err)
	}
	log.Println("Reading from file, ", err)
	recodrs, err = reader.ReadAll()
	fmt.Println(recodrs)

	writer.Flush()
	
	log.Println("End of Create function... ", err)
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
