package csv

import (
	"encoding/csv"
	"imersao16-ordenation/internal/dto"
	"os"
	"strconv"
)

type CSVTransform struct {
}

func NewCsvTransform() CSVTransform {
	return CSVTransform{}
}

func (c CSVTransform) Read(filename string) ([]dto.Person, error) {
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

	var people []dto.Person
	for _, line := range lines {
		age, _ := strconv.Atoi(line[1])
		score, _ := strconv.Atoi(line[2])
		person := dto.Person{
			Name:  line[0],
			Age:   age,
			Score: score,
		}
		people = append(people, person)
	}

	return people, nil
}

func (c CSVTransform) Write(filename string, people []dto.Person) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, person := range people {
		record := []string{person.Name, strconv.Itoa(person.Age), strconv.Itoa(person.Score)}
		err := writer.Write(record)
		if err != nil {
			return err
		}
	}

	return nil
}
