package struct_demo

import (
	"fmt"
)

type Person struct {
	firstName string
	lastName string
}

func StructDemo() {
	p1 :=Person{}
	p1.firstName = "Will"
	p1.lastName = "Smith"
	fmt.Println(p1)

	p2 :=Person{}

	fmt.Println(p2)
}