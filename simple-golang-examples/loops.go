package main

func forLoops() {
	var i int
	//simple conditional loop
	println("\n conditional loop -> for i < 5, break on 3")
	for i < 5 {
		println(i)
		i++
		if i == 3 {
			break
		}
	}

	var j int
	println("\n conditional loop -> for j < 5, continue on 3")
	println("\n j value = ", j)
	for j < 5 {
		println(j)
		j++
		if j == 3 {
			continue
		}
		println("counting...")
	}

	// standard for loop -> for(int i; i < len; i++)
	for z := 0; z < 5; z++ {
		println("printing", z)
	}

	// infinite loop that will go forever until break occurs
	var q int
	// other way is for ; ; {}
	for {
		if q > 5 {
			break
		}
		println("inf loop count", q)
		q++
	}

}

func collectionLoops() {
	slice := []int{1, 2, 3}
	//post condition clause for loop
	for i := 0; i < len(slice); i++ {
		println(slice[i])
	}

	// value condition
	for i, v := range slice {
		println("index = ", i, " & value = ", v)
	}

	m := map[string]int{"ReqLimit": 13, "port": 3000}
	for k, val := range m {
		println("key = ", k, " & val = ", val)
	}

	// only keys
	println("\n print only keys")
	for k := range m {
		println("key = ", k)
	}

	// only keys
	println("\n print only values")
	for _, v := range m {
		println("vals = ", v)
	}

}
