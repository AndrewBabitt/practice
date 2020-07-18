package main

import (
	"fmt"
)

func collections() {
	fmt.Println(" \n \n ==== Collections ==== ")
	fmt.Println("\n ==== Slices ==== ")
	fmt.Println(" built ontop of array, same constructs, dyanmically sized ")
	//using an array to build a slice
	arr1 := [3]int{7, 8, 9}
	//slice data type of all the elements of an array
	slicearr := arr1[:]

	//arr and slicearr ar ethe same
	fmt.Println("arr1, slicearr", arr1, slicearr)
	arr1[1] = 42
	slicearr[2] = 26
	//slice is pointing to underlying array so it will be updated from initial array and slice value
	fmt.Println(arr1, slicearr)

	//creating a slice standalone
	slice := []int{1, 2, 3}
	fmt.Println(slice)

	slice2 := []int{} //empty slice
	fmt.Println(slice2)
	fmt.Println("Append single value to list")
	slice2 = append(slice2, 2) // add single elemnt
	fmt.Println(slice2)
	fmt.Println("appending multiple values to list")
	slice2 = append(slice2, 4, 5, 6, 7, 8) // add multiple elements
	fmt.Println(slice2)

	left := slice2[:3]
	right := slice2[4:]
	fmt.Println("left := slice2[:3]")
	fmt.Println("right := slice2[4:]")
	fmt.Println("left = ", left)
	fmt.Println("right = ", right)

	fmt.Println("\n ==== Arrays ==== ")
	fmt.Println(" fixed size collection of similar data types ")
	var arr [3]int               //explicit createing an array with 3 elements of ints
	arrInline := [3]int{4, 5, 6} // short hand
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr)
	fmt.Println("arr[2] = ", arr[2])
	// fmt.Println("out of bounds => arr[4]", arr[4]) // arr[4] => invalid array index 4 (out of bounds for 3-element array)
	fmt.Println("arrInline => ", arrInline)

	fmt.Println("\n ==== Maps ==== ")
	fmt.Println(" Collection types with Key value pairs ")
	fmt.Println(" shorthand ")
	fmt.Println(" varName := map[keyType]valueType ")
	//short hand

	m := map[string]int{"foo": 3, "bar": 2, "baz": 39}
	fmt.Println("map m output", m)
	fmt.Println("m[\"foo\"] value = ", m["foo"])

	fmt.Println("delete func = delete(mapName, key)")
	fmt.Println("delete baz ")
	delete(m, "baz")
	fmt.Println(m)

	fmt.Println("\n ==== Structs ====")
	fmt.Println(" TBD ")
}
