package main

import "fmt"

type myStructure struct {
	Name    string
	Surname string
	Height  int32
}

func NewStructurePointer(n, s string, h int32) *myStructure {
	if h > 300 {
		h = 0
	}
	return &myStructure{n, s, h}
}

func NewStructure(n, s string, h int32) myStructure {
	if h > 300 {
		h = 0
	}
	return myStructure{n, s, h}
}

func main() {
	s1 := NewStructurePointer("Andrey", "Solodov", 182)
	s2 := NewStructure("Andrey", "Solodov", 182)

	fmt.Println((*s1).Name)
	fmt.Println(s2.Name)
	fmt.Println(s1)
	fmt.Println(s2)

	pS := new(myStructure)
	fmt.Println(pS)
	fmt.Printf("%T\n", pS)
}
