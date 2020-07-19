package main

import (
	"errors"
	"fmt"
)

func tryingOutFunction() {
	fmt.Println("starting this func")
	//do more things
	fmt.Println("ending this func")
}

func paramFunction(p1 int) {
	fmt.Println("== param Function ==")
	fmt.Println("p1 value = ", p1)
}

func multiplyInts(pm1 int, pm2 int) int {
	return pm1 * pm2
}

func dealingWithErrors(p string) (string, error) {
	if p == "ha" {
		return "ya!", nil
	}

	return "", errors.New("something went wrong")
}
