package main

import "fmt"

func primitives() {
	// verbose declaration
	var i int
	i = 42
	fmt.Println(i)

	// inline
	var f float32 = 3.14
	fmt.Println(f)

	//common var delcaration
	// type is implied by compiler
	name := "Mike Jones"
	fmt.Println(name)

	//THIS WILL THROW A COMPILER ERRO SINCE IT'S NEVER USED
	//namer := "Jones" REMOVED

	// boolean
	b := true
	fmt.Println(b)

	// complex datatypes built in
	c := complex(3, 4)
	fmt.Println(c)

	//multi assignment with complex types
	r, im := real(c), imag(c)
	fmt.Println(r, im)

	// Pointers (*)
	var fName *string
	fmt.Println(fName) //nil until pointing

	//*fName = "John"    // deference operator
	//fmt.Println(fName) //uninitialized reference == nil

	// go wants to make sure there is an explicit location in memory otherwise it will complain

	var firstName *string = new(string) //initializes pointer to mem ref

	*firstName = "Mike"
	fmt.Println(firstName)  //memory reference
	fmt.Println(*firstName) // value pointing to

	// no pointer arithmatic

	//Address Of (&)
	newName := "Andrew"

	ptr := &newName
	fmt.Println(ptr, *ptr) // memory address, pointer value
}

func constants() {

	const pi = 3.1415 //constants and cannot be changed and needs to evaluated at compile time
	fmt.Println(pi)
	// pi = 3.1 // can't be changes

	// implictly type constant
	const x = 3
	fmt.Println(x + 3) //x is interpreted as int by compiler

	// ...lines of code...
	fmt.Println(x + 2.5) //x is intepreted as float by compiler

	// explict type constant
	const z int = 3
	fmt.Println(z + 3) // x is an int, so valid

	// ...lines of code...
	fmt.Println(float32(z) + 2.5) //converted to float32 num to compile

}
