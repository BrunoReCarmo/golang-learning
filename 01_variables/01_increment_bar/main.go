package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	//number of cols
	const col = 30
	//bar var
	bar := fmt.Sprintf("\x0c[%%-%vs]", col)

	//D will in at O
	//While D is less than col(30) will increment
	for d := 0; d < col; d++ {
		fmt.Printf(bar, strings.Repeat("=", d)+">")
		//Sleep - (Break between increment)
		time.Sleep(100 * time.Millisecond)
	}
	//After incremeent is over, the output will be Done sir! ;-)
	// string.Repeat returns a new string consisting of the input string repeated count times
	fmt.Printf(bar+"Done sir! ;-)", strings.Repeat("=", col))
}
