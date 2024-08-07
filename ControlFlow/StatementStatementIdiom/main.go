package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("This is the first statement to run")
	fmt.Println("This is the second statement to run")
	x := 40 // this is the third statement to run
	y := 5  // this is the fourth statement to run
	fmt.Printf("x = %v \ny = %v\n", x, y)

	/*
		the expression [evaluated in an if statement] may be preceded by a simple statement, which executes before the expression is evaluated

		the comma ok idiom is also like this
	*/

	if z := 2 * rand.Intn(x); z >= x {
		fmt.Printf("z is %v and that is GREATER THAN OR EQUAL x which is %v\n", z, x)
	} else {
		fmt.Printf("z is %v and that is LESS THAN x which is %v\n", z, x)
	}
}
