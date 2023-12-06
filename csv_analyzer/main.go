package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
    // Open the CSV file
    file, err := os.Open("example.csv")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    // Create a new reader
    reader := csv.NewReader(file)

    // Read all records
    records, err := reader.ReadAll()
    if err != nil {
        fmt.Println("Error reading CSV data:", err)
        return
    }

    // Process records
    for _, record := range records {
        fmt.Println(record) // Each record is a slice of fields
    }
}
