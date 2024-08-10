package createfile

import (
	"log"
	"os"
	"encoding/csv"
)

func CreateFile() {
	// Check if the data.csv file exists.
	if fileExists("data.csv") {
		log.Println("File is already there")
	} else {
		file, err := os.Create("data.csv")
		if err != nil {
			log.Fatal(err)
		}
		writeHeader()
		defer file.Close()
	}
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}


func writeHeader(){
	// Create a new CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write the header to the CSV file
	heading := []string{"Id", "Heading", "Description", "Duration"}
	err = writer.Write(heading)
	if err != nil {
		log.Fatal(err)
	}
}