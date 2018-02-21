package main

import "fmt"

func main() {
	// loop from i := 0 to i < 9
	fmt.Println("for i:=0 is not defined while i is less than 10")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// we can use for loop like while loop in following way
	fmt.Println("for like while")
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
	// infinite loop
	/*
		for {
			fmt.Println("Inside infinite loop")
		}
	*/

	// breaking out of loop
	x := 5
	for { // infinite loop
		fmt.Println("breaking loop ", x)
		x += 3 // x + 3
		if x > 25 {
			break // Breaking when condition reached
		}
	}
}
