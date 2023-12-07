package csv

import (
	"encoding/csv"
	"fmt"
	"imersao16-ordenation/internal/dto"
	"os"
	"sort"
)

type CSVTransform struct {
}

func NewCsvTransform() CSVTransform {
	return CSVTransform{}
}

func (c CSVTransform) Read(filename string) ([]dto.CsvFile, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var people []dto.CsvFile
	for i, line := range lines {
		var CsvFile dto.CsvFile
		if i == 0 {
			CsvFile.Header = line
		}

		CsvFile = dto.CsvFile{
			Line: line,
		}

		people = append(people, CsvFile)
	}

	fmt.Println(people)
	return people, nil
}

func (c CSVTransform) Write(filename string, people []dto.CsvFile) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, record := range people {

		if record.Header != nil {
			err := writer.Write(record.Header)
			if err != nil {
				return err
			}
		}

		err := writer.Write(record.Line)
		if err != nil {
			return err
		}

	}

	return nil
}

func (c CSVTransform) SortedByAge(people []dto.CsvFile, outputFileName string) error {
	sort.Sort(dto.ByAge(people))

	err := c.Write(outputFileName, people)
	if err != nil {
		fmt.Printf("Error writing output file sorted by age: %s\n", err)
		return err
	}

	return nil
}

func (c CSVTransform) SortedByName(people []dto.CsvFile, outputFileName string) error {
	sort.Sort(dto.ByName(people))

	err := c.Write(outputFileName, people)
	if err != nil {
		fmt.Printf("Error writing output file sorted by name: %s\n", err)
		return err
	}

	return nil
}
