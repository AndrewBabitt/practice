package main

import "fmt"

type Example struct {
	ID   int
	Name string
	Desc string
}

func ifStatements() {
	e1 := Example{
		ID:   5,
		Name: "John Bravo",
		Desc: "Oh yeah!",
	}
	e2 := Example{
		ID:   7,
		Name: "Happy Gilmore",
		Desc: "Tap Tap!",
	}
	fmt.Println(e1)
	fmt.Println(e2)

	if e1.ID == e2.ID {
		println("users are the same")
	} else {
		println("users are not the same")
	}

	if e1 == e2 {
		println("these are the same")
	} else if e1.Desc == e2.Desc {
		println("these 2 do the same thing")
	} else {
		println("they are not the same")
	}

}
