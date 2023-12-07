package dto

import "strings"

type Person struct {
	Name  string
	Age   int
	Score int
}

type ByName []Person
type ByAge []Person

func (a ByName) Len() int      { return len(a) }
func (a ByName) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool {
	namei := strings.ToUpper(a[i].Name)
	nameJ := strings.ToUpper(a[j].Name)
	return namei < nameJ
}

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
