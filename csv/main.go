package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

type Employee struct {
	ID  string
	Age int
}

func main() {
	records := []Employee{
		{"E01", 25},
		{"E02", 26},
		{"E03", 24},
		{"E04", 26},
	}
	file, err := os.Create("records.csv")
	defer file.Close()
	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	w := csv.NewWriter(file)
	defer w.Flush()
	// Using Write
	for _, record := range records {
		row := []string{record.ID, strconv.Itoa(record.Age)}
		if err := w.Write(row); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}

	// Using WriteAll
	var data [][]string
	for _, record := range records {
		row := []string{record.ID, strconv.Itoa(record.Age)}
		data = append(data, row)
	}
	w.WriteAll(data)
}
