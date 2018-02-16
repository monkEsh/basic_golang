package main

// Import packages needed
// fmt is use format to do basic io operations
import (
	"fmt"
)

// function declaration is done in below format
// func <function_name>(<parameters><type>, <parameters><type>)(<return datatype>, <return datatype>){
// operations
// return <var1>, <var2>...
//}
func multiple(a, b string) (string, string) {
	return a, b
}

func add(x, y float64) float64 {
	return x + y
}

func main() {
	w1, w2 := "hello", "World"
	num1, num2 := 5.6, 9.5
	fmt.Println(multiple(w1, w2))
	fmt.Println(add(num1, num2))
}
