package csvparser

import (
	"os"
	"testing"
	"trip-route/services/utils"
)

type TestDataItem struct {
	inputPath  string
	outputData [][]string
}

func TestReadCsv(t *testing.T) {
	pathFile1 := "testdata/file1.csv"
	outputFile1 := [][]string{
		{"GRU", "BRC", "20"},
		{"BRC", "SCL", "5"},
	}
	pathFile2 := "testdata/file2.csv"
	outputFile2 := [][]string{
		{"GRU", "CDG", "75"},
		{"SCL", "ORL", "20"},
	}

	dataItems := []TestDataItem{
		{pathFile1, outputFile1},
		{pathFile2, outputFile2},
	}

	for _, item := range dataItems {
		err, result := Read(item.inputPath)

		if err != nil {
			t.Errorf("csvparser.Read() from path %v FAILED, expected %v but got an error '%v'", item.inputPath, item.outputData, err)
		} else if len(result) != 2 {
			t.Errorf("csvparser.Read() from path %v FAILED, expected %v but got value '%v'", item.inputPath, item.outputData, result)
		} else {
			t.Logf("csvparser.Read() from path %v PASSED, expected %v and got '%v'", item.inputPath, item.outputData, result)
		}
	}
}

func BeforeTestWriteCsv(filePaths []string) []string {
	result := make([]string, len(filePaths))

	// Copy csv files only for this test that appends rows to the end of file
	for index, item := range filePaths {
		newPathName := utils.FilenameTrimmedSuffix(item) + "_testwrite.csv"
		err := utils.Copy(item, newPathName)
		if err != nil {
			// Todo: improve error handling in this case
			result[index] = item
		} else {
			result[index] = newPathName
		}
	}

	return result
}

func AfterTestWriteCsv(filePaths []string) {
	for _, item := range filePaths {
		os.Remove(item)
	}
}

func TestWriteCsv(t *testing.T) {
	inputPathFile1 := "testdata/file1.csv"
	outputFile1 := [][]string{
		{"GRU", "BRC", "20"},
		{"BRC", "SCL", "5"},
		{"GRU", "BRC", "20"},
		{"BRC", "SCL", "5"},
	}
	inputPathFile2 := "testdata/file2.csv"
	outputFile2 := [][]string{
		{"GRU", "CDG", "75"},
		{"SCL", "ORL", "20"},
		{"GRU", "CDG", "75"},
		{"SCL", "ORL", "20"},
	}

	outputPaths := BeforeTestWriteCsv([]string{inputPathFile1, inputPathFile2})

	dataReadItems := []TestDataItem{
		{inputPathFile1, outputFile1},
		{inputPathFile2, outputFile2},
	}
	dataWriteItems := []TestDataItem{
		{outputPaths[0], outputFile1},
		{outputPaths[1], outputFile2},
	}

	for index, item := range dataReadItems {
		errRead, readResult := Read(item.inputPath)

		if errRead != nil {
			t.Errorf("csvparser.Write() from path %v FAILED, Read() wasn't successful, then it was expected %v but got an error '%v'", item.inputPath, item.outputData, errRead)
		} else {
			errWrite, writeResult := Write(dataWriteItems[index].inputPath, readResult)
			if errWrite != nil {
				t.Errorf("csvparser.Write() from path %v FAILED, expected %v but got error '%v'", dataWriteItems[index], item.outputData, errWrite)
			} else if len(writeResult) != 4 {
				t.Errorf("csvparser.Write() from path %v FAILED, expected %v but got value '%v'", dataWriteItems[index], item.outputData, writeResult)
			} else {
				t.Logf("csvparser.Write() from path %v PASSED, expected %v and got '%v'", dataWriteItems[index], item.outputData, writeResult)
			}
		}
	}

	AfterTestWriteCsv(outputPaths)
}

func TestCreateWrite(t *testing.T) {
	inputPathFile1 := "testdata/createwrite_test1.csv"
	outputFile1 := [][]string{
		{"GRU", "BRC", "20"},
		{"BRC", "SCL", "5"},
		{"GRU", "BRC", "20"},
		{"BRC", "SCL", "5"},
	}
	inputPathFile2 := "testdata/createwrite_test1.csv"
	outputFile2 := [][]string{
		{"GRU", "CDG", "75"},
		{"SCL", "ORL", "20"},
		{"GRU", "CDG", "75"},
		{"SCL", "ORL", "20"},
	}

	dataWriteItems := []TestDataItem{
		{inputPathFile1, outputFile1},
		{inputPathFile2, outputFile2},
	}

	for _, item := range dataWriteItems {
		err, result := CreateWrite(item.inputPath, item.outputData)

		if err != nil {
			t.Errorf("csvparser.CreateWrite() from path %v FAILED, it was expected %v but got an error '%v'", item.inputPath, item.outputData, err)
		} else if len(result) != 4 {
			t.Errorf("csvparser.CreateWrite() from path %v FAILED, expected %v but got value '%v'", item.inputPath, item.outputData, result)
		} else {
			t.Logf("csvparser.CreateWrite() from path %v PASSED, expected %v and got '%v'", item.inputPath, item.outputData, result)
		}
	}

	AfterTestWriteCsv([]string{inputPathFile1, inputPathFile2})
}
