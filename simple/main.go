package main

import (
	"fmt"
)

func main() {

	// sanity check that it compiles
	fmt.Println("Hello, Everyone!")
	primitives()
	constants()
	g()
	collections()
	t := tryingOutFunction // treated as a function as an object that can be passed around the app
	t()                    // invoking function assigned to t
	tryingOutFunction()    // invokes the functions
	paramFunction(42)
	pf := paramFunction
	pf(100)
	product := multiplyInts(5, 10)
	fmt.Println(product)
	r, err := dealingWithErrors("ha")
	fmt.Println("retrun value from dealing with errors = ", r)
	fmt.Println("error is ? ", err)
	r1, err1 := dealingWithErrors("nah")
	fmt.Println("retrun value from dealing with errors = ", r1)
	fmt.Println("error is ? ", err1)

	//if there's an error that you only care about and return value is not worried about
	_, err2 := dealingWithErrors("")
	if err2 != nil {
		fmt.Println("Has error =>", err2)
	}
	_, err3 := dealingWithErrors("ha")
	if err3 != nil {
		fmt.Println("Has error =>", err2)
	} else {
		fmt.Println("never hit an error")
	}
}
