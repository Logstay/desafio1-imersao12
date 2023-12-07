package main

import (
	"fmt"
	"imersao16-ordenation/internal/dto"
	csv "imersao16-ordenation/internal/transform"
	"os"
	"sort"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go input.csv output.csv")
		return
	}

	transform := csv.NewCsvTransform()

	inputFilename := os.Args[1]
	outputFilename := os.Args[2]

	people, err := transform.Read(inputFilename)
	if err != nil {
		fmt.Printf("Error reading input file: %s\n", err)
		return
	}

	// Sort by name
	sort.Sort(dto.ByName(people))
	err = transform.Write(outputFilename+"_sorted_by_name.csv", people)
	if err != nil {
		fmt.Printf("Error writing output file sorted by name: %s\n", err)
		return
	}

	// Sort by age
	sort.Sort(dto.ByAge(people))
	err = transform.Write(outputFilename+"_sorted_by_age.csv", people)
	if err != nil {
		fmt.Printf("Error writing output file sorted by age: %s\n", err)
		return
	}

	fmt.Println("Files created successfully.")
}
