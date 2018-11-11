package main

import (
	"fmt"
)

func main() {
	complexArray1 := [3][]string {
		[]string{"d", "e", "f"},
		[]string{"g", "h", "i"},
		[]string{"j", "k", "l"},
	}

	b := modifyValue(complexArray1)
	fmt.Printf("The modified value array: %v\n", b)
	fmt.Printf("The original complex array: %v\n", complexArray1)

	fmt.Println("----------------------\n")

	complexArray2 := [3][]string {
		[]string{"d", "e", "f"},
		[]string{"g", "h", "i"},
		[]string{"j", "k", "l"},
	}
	c := modifyReference(complexArray2)
	fmt.Printf("The modified reference array address: %p\n", &c)
	fmt.Printf("The modified value array: %v\n", c)
	fmt.Printf("The original complex array: %v\n", complexArray2)
}

func modifyValue(a [3][]string) [3][]string {
	a[1] = []string{"x", "y", "x"}
	return a
}

func modifyReference(a [3][]string) [3][]string {
	a[1][1] = "O"

	fmt.Printf("The address of a: %p\n", &a)
	return a
}