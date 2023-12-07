package dto

import "strconv"

type CsvFile struct {
	Header []string
	Line   []string
}

type ByName []CsvFile
type ByAge []CsvFile

func (a ByName) Len() int      { return len(a) }
func (a ByName) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool {
	// Assuming the 'Name' column is at index 0 in the 'Line' field
	return a[i].Line[0] < a[j].Line[0]
}

func (a ByAge) Len() int      { return len(a) }
func (a ByAge) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool {
	// Assuming the 'Age' column is at index 1 in the 'Line' field
	ageA, _ := strconv.Atoi(a[i].Line[1])
	ageB, _ := strconv.Atoi(a[j].Line[1])
	return ageA < ageB
}
