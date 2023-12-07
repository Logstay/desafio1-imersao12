package main

import (
	"fmt"
	csv "imersao16-ordenation/internal/transform"
	"os"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Println("Please provide 3 input filenames. 1:")
		return
	}

	transform := csv.NewCsvTransform()

	inputFileName := os.Args[1]

	people, err := transform.Read(inputFileName)
	if err != nil {
		fmt.Printf("Error reading input file: %s\n", err)
		return
	}

	outputFileName := os.Args[2]
	err = transform.SortedByName(people, outputFileName)
	if err != nil {
		fmt.Printf("Error sorting by name: %s\n", err)
		return
	}

	outputSecondFileName := os.Args[3]
	err = transform.SortedByAge(people, outputSecondFileName)
	if err != nil {
		fmt.Printf("Error sorting by age: %s\n", err)
		return
	}

	fmt.Println("Files created successfully.")
}
