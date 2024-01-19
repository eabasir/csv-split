package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatalf("Usage: go run main.go <csv_filename> <number_of_records_in_each_file>")
	}
	csvFilename := os.Args[1]

	numRecordsStr := os.Args[2]
	numRecords, err := strconv.Atoi(numRecordsStr)
	if err != nil {
		log.Fatalf("Failed to parse the number of records: %v", err)
	}

	outputDir := filepath.Join(filepath.Dir(os.Args[0]), "outputs")
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		log.Fatalf("Failed to create the output directory: %v", err)
	}

	inputFile, err := os.Open(csvFilename)
	if err != nil {
		log.Fatalf("Failed to open the input CSV file: %v", err)
	}
	defer inputFile.Close()

	reader := csv.NewReader(bufio.NewReader(inputFile))

	header, err := reader.Read()
	if err != nil {
		log.Fatalf("Failed to read the header from the input file: %v", err)
	}

	fileIndex := 1
	recordCount := 0
	outputFile, err := createOutputFile(outputDir, fileIndex)
	if err != nil {
		log.Fatalf("Failed to create output file: %v", err)
	}
	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	writer.Write(header)

	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Failed to read a row from the input file: %v", err)
		}

		recordCount++
		writer.Write(row)

		if recordCount == numRecords {
			outputFile.Close()

			fileIndex++
			outputFile, err = createOutputFile(outputDir, fileIndex)
			if err != nil {
				log.Fatalf("Failed to create output file: %v", err)
			}
			writer = csv.NewWriter(outputFile)
			defer writer.Flush()

			writer.Write(header)

			recordCount = 0
		}
	}

	fmt.Printf("Split complete! Total number of split files: %d\n", fileIndex)
}

func createOutputFile(outputDir string, fileIndex int) (*os.File, error) {
	filename := fmt.Sprintf("output_%d.csv", fileIndex)
	filePath := filepath.Join(outputDir, filename)
	return os.Create(filePath)
}
