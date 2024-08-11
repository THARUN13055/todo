package checkenv

import (
	"encoding/csv"
	"log"
	"os"
)

func CreateFile() {
	// Check if the data.csv file exists.
	filenameed := "data.csv"
	if !fileExists(filenameed) {
		file, err := os.Create(filenameed)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		writeHeading(file)
	}

}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func writeHeading(file *os.File) {
	// Create a new CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write the header to the CSV file
	heading := []string{"Id", "Heading", "Description", "Duration"}
	err := writer.Write(heading)
	if err != nil {
		log.Fatal(err)
	}
}
