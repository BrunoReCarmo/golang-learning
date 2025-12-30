package main

import "fmt"

func main() {
	x := 2
	y := 6

	plus := x + y
	minus := y-x
	times := x*y 
	dividedBy := y/x
 
	//Rule of three -> if y == 100%, x will be 1/3 or 33%
	ruleOf3 := (x * 100) / y

	const formatOutput string = `
	{
		"plus": %d,
		"minus": %d,
		"times": %d,
		"divided by": %d,
		"Rule of 3": %d,
	}`
	//Console:
	// {
	// 	"plus": 8,
	// 	"minus": 4,
	// 	"times": 12,
	// 	"divided by": 3,
	// 	"Rule of 3": 33,
	// }
	fmt.Printf(formatOutput, plus, minus, times, dividedBy, ruleOf3)
}