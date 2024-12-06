package main

import "fmt"

type Student struct {
	FirstName string
	LastName  string
}

type Class struct {
	Students []Student
}

func main() {
	c1 := Class{[]Student{{"Diego", "maroni"}, {"Arber", "Schneider"}, {"Lara", "Weber"}}}
	c2 := Class{[]Student{{"Max", "Koch"}, {"Michael", "Lehmann"}, {"Paul", "Schmidt"}}}

	m := map[int][]Class{
		101: {c1},
		202: {c2},
		303: {c1, c2},
	}

	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(m)
}
