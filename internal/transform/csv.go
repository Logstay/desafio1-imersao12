package csv

import (
	"encoding/csv"
	"fmt"
	"imersao16-ordenation/internal/dto"
	"os"
	"sort"
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
	for i, line := range lines {
		if i == 0 {
			continue
		}

		age, _ := strconv.Atoi(line[1])
		score, _ := strconv.Atoi(line[2])
		person := dto.Person{
			Name:  line[0],
			Age:   age,
			Score: score,
		}
		people = append(people, person)
	}

	fmt.Println(people)
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

	data := [][]string{
		{"Nome", "Idade", "Pontuação"},
	}

	for _, record := range people {
		data = append(data, []string{record.Name, strconv.Itoa(record.Age), strconv.Itoa(record.Score)})
	}

	err = writer.WriteAll(data)
	if err != nil {
		fmt.Println("Writing file error write...")
		return err
	}

	return nil
}

func (c CSVTransform) Sort(people []dto.Person, outputFileName string) error {
	sort.Sort(dto.ByName(people))

	err := c.Write(outputFileName, people)
	if err != nil {
		fmt.Printf("Error writing output file sorted by name: %s\n", err)
		return err
	}

	return nil
}
