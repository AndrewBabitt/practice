package main

import (
	"fmt"
)

const pi = 3.1415

// const (
// 	// allows for evolution of a constant
// 	first  = iota + 6
// 	second = 2 << iota
// )

const (
	first = iota
	second
	third
)

//const blocks that are on the top
const (
	forth = iota
)

func g() {
	fmt.Println("first, second, third")
	fmt.Println(first, second, third)
	fmt.Println("forth", forth) // resets to 0
}
