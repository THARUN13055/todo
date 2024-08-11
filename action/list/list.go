package list

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func List() {
	f,err := os.Open("data.csv")
	if err != nil {
		log.Fatal("The File has not been open",err)
	}
	defer f.Close()

	csvFileReader := csv.NewReader(f)

	//Reading the Header

	header , err := csvFileReader.Read()
	if err != nil {
		log.Fatal("Error reading CSV header: ", err)
	}

	fmt.Println(header)

	if len(os.Args) > 2 && (os.Args[2] == "--key" || os.Args[2] == "-k") {
		value := os.Args[3]

		if value == "" {
			value = "Null"
		}

		for {
			row , err := csvFileReader.Read()
			if err != nil {
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

	for {
		row , err := csvFileReader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			log.Fatal("Error reading CSV row: ", err)
		}
		fmt.Println(row)

	}
}