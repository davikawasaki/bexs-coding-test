package csvparser

import (
	"encoding/csv"
	"errors"
	"os"
	"trip-route/services/utils"
)

func Read(path string) (error, [][]string) {
	// Takes a string with the file path and returns a file descriptor to open the file
	csvfile, err := os.Open(path)
	if err != nil {
		return err, nil
	}

	// Parse the file
	r := csv.NewReader(csvfile)

	// Read all records at once
	records, _ := r.ReadAll()
	return nil, records
}

func CreateWrite(path string, csvData [][]string) (error, [][]string) {
	if csvData == nil {
		return errors.New("No data to be written."), nil
	}

	if path == "" {
		return errors.New("Couldn't open an empty path"), nil
	}

	csvfile, err := os.Create(path)
	if err != nil {
		return errors.New("File couldn't be created to be written!"), nil
	}

	w := csv.NewWriter(csvfile)

	err = w.WriteAll(csvData)
	if err != nil {
		return err, nil
	}

	return nil, csvData
}

func Write(path string, csvData [][]string) (error, [][]string) {
	if csvData == nil {
		return errors.New("No data to be written."), nil
	}

	if path == "" {
		return errors.New("Couldn't open an empty path"), nil
	}

	dataVals := [][]string{}
	var dataFile [][]string

	if !utils.FileExists(path) {
		_, err := os.Create(path)
		if err != nil {
			return errors.New("Non-existent file couldn't be created to be written!"), nil
		}
	} else {
		err, data := Read(path)
		if err != nil {
			return errors.New("Existent file couldn't be read to be written!"), nil
		}
		dataFile = data
	}

	if dataFile == nil {
		dataFile = dataVals
	} else {
		for _, item := range csvData {
			dataFile = append(dataFile, item)
		}
	}

	csvfile, err := os.Create(path)
	if err != nil {
		return err, nil
	}

	w := csv.NewWriter(csvfile)

	err = w.WriteAll(dataFile)
	if err != nil {
		return err, nil
	}

	return nil, dataFile
}
