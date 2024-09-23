package main

import (
	"io"
	"os"

	"github.com/gocarina/gocsv"
)

func (app App) ReadCsv() []*Issues {
	// Try to open the example.csv file in read-write mode.
	csvFile, csvFileError := os.OpenFile(app.CsvFile, os.O_RDWR, os.ModePerm)
	// If an error occurs during os.OpenFIle, panic and halt execution.
	if csvFileError != nil {
		panic(csvFileError)
	}
	// Ensure the file is closed once the function returns
	defer csvFile.Close()

	var issues []*Issues

	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		return gocsv.LazyCSVReader(in)
	})
	// Parse the CSV data into the articles slice. If an error occurs, panic.
	if unmarshalError := gocsv.UnmarshalFile(csvFile, &issues); unmarshalError != nil {
		panic(unmarshalError)
	}

	return issues
}
