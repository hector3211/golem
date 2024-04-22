package csv

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Lot struct {
	Name     string
	Location string
	Price    int
	Status   string
}

func check(err error) {
	if err != nil {
		log.Fatal("Error:", err)
		return
	}
}

func GetFile(filepath string) [][]string {
	file, err := os.Open(filepath)

	// Checks for the error
	check(err)

	// Closes the file
	defer file.Close()

	// The csv.NewReader() function is called in
	// which the object os.File passed as its parameter
	// and this creates a new csv.Reader that reads
	// from the file
	reader := csv.NewReader(file)

	// ReadAll reads all the records from the CSV file
	// and Returns them as slice of slices of string
	// and an error if any
	record, err := reader.ReadAll()

	// Checks for the error
	check(err)

	return record[1:]
}

func ReadFile(filePath string) {
	// os.Open() opens specific file in
	// read-only mode and this return
	// a pointer of type os.File
	file, err := os.Open(filePath)

	// Checks for the error
	check(err)

	// Closes the file
	defer file.Close()

	// The csv.NewReader() function is called in
	// which the object os.File passed as its parameter
	// and this creates a new csv.Reader that reads
	// from the file
	reader := csv.NewReader(file)

	// ReadAll reads all the records from the CSV file
	// and Returns them as slice of slices of string
	// and an error if any
	records, err := reader.ReadAll()

	// Checks for the error
	check(err)

	// Loop to iterate through
	// and print each of the string slice
	var inventory []Lot
	for _, record := range records[1:] {
		price, err := strconv.Atoi(record[2])
		check(err)
		lot := Lot{
			Name:     record[0],
			Location: record[1],
			Price:    price,
			Status:   record[3],
		}
		inventory = append(inventory, lot)
	}
	for _, slab := range inventory {

		fmt.Printf("Name: %s Location: %s Price_per_sqft: %d Status: %s\n", slab.Name, slab.Location, slab.Price, slab.Status)
	}
}
