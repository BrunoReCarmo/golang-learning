package main

import "fmt"

// isEqual compares two integers and returns true if they are equal
func isEqual(a int, b int) bool {
	return a == b
}

func main() {
	y := 4
	x := 5
	z := 10 - 6

	fmt.Println(isEqual(y, x)) // false
	fmt.Println(isEqual(y, z)) // true
	fmt.Println(isEqual(x, z)) // false
}
