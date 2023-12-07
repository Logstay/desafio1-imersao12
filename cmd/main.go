package main

import (
	"fmt"
	csv "imersao16-ordenation/internal/transform"
	"os"
)

func main() {

	transform := csv.NewCsvTransform()

	inputFileName := os.Args[1]

	people, err := transform.Read(inputFileName)
	if err != nil {
		fmt.Printf("Error reading input file: %s\n", err)
		return
	}

	outputFileName := os.Args[2]
	err = transform.Sort(people, outputFileName)
	if err != nil {
		fmt.Printf("Error sorting: %s\n", err)
		return
	}

	fmt.Println("Files created successfully.")
}
