package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <csv_file_path> <split_size>")
		return
	}

	filePath := os.Args[1]
	splitSize, err := strconv.Atoi(os.Args[2])
	if err != nil || splitSize <= 0 {
		fmt.Println("Invalid split size. It must be a positive integer.")
		return
	}

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		fmt.Printf("Error reading CSV: %v\n", err)
		return
	}

	if len(rows) < 2 {
		fmt.Println("CSV file must have at least a header and one data row.")
		return
	}

	header := rows[0]
	dataRows := rows[1:]

	for i, start := 0, 0; start < len(dataRows); i, start = i+1, start+splitSize {
		end := start + splitSize
		if end > len(dataRows) {
			end = len(dataRows)
		}

		outputFileName := fmt.Sprintf("output_%d.csv", i+1)
		outputFile, err := os.Create(outputFileName)
		if err != nil {
			fmt.Printf("Error creating file %s: %v\n", outputFileName, err)
			return
		}
		defer outputFile.Close()

		writer := csv.NewWriter(outputFile)
		defer writer.Flush()

		writer.Write(header)
		writer.WriteAll(dataRows[start:end])
	}

	fmt.Println("CSV split completed successfully.")
}
