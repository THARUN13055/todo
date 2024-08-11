package list

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func list() {
	value := os.Args[3]
	if os.Args[1] == "--list" || os.Args[1] == "-t" {
		f,err := os.Open("data.csv")
		err != nil {
			log.Println("The File has not been open",err)
		}
		defer f.Close()
	
		csvFileReader := csv.NewReader(f)
		readallfile,err := csvFileReader.ReadAll()
		if err != nil {
			log.Fatal("Unable to parse file as CSV for ", err)
		}
	
		fmt.Println(csvFileReader)
	}

	if os.Args[2] == "--key" || os.Args == "-k" {
		f,err := os.Open("data.csv")
		err != nil {
			log.Fatal("The  error is in opening file specific read:->",err)
		}
		defer f.Close()
	
		csvFileReader := csv.NewReader(f)

		//Reading the Header

		header , err := csvFileReader.Read()
		if err != nil {
			log.Fatal("Error reading CSV header: ", err)
		}

		fmt.Println(header)

		for {
			row , err := csvFileReader.Read()
			err != nil {
				if err.Error() == "EOF" {
					break
				}
				log.Fatal("Error reading CSV row: ", err)
			}
			if row[3] == value {
				fmt.Println(row)
			}

		} 
	}

}