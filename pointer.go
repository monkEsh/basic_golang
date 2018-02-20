// Every Go program is made up of packages.

// Programs start running in package main.

package main

// Package fmt implements formatted I/O with functions analogous to C's printf and scanf.
import (
	"fmt"
)

func main() {
	// variable declared
	x := 15

	// storing memory address of x in a
	a := &x

	fmt.Println("Memory address : ", a)

	// printing value of value at memory location a by *
	fmt.Println("Value stored at a ", *a)

	// assigning new value to variable using memory loction
	*a = 5

	// print x
	fmt.Println("Value of x after chaanging value at memory location ", x)

}
