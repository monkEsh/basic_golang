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
	/*
		we can initialize variables by following methods
		1. Declare variables with datatype
			var <variable name> <datatype> = <value>
			var <variable name> <datatype> = <value>
			var <variable name> <datatype> = <value>
		2. Declare multiple variables of same type in line
			var <var name1>, <var num2> <datatype> = <var1 value>, <var2 value>
		3. Declare variables inside function without datatype, Go dev use this method to define variable inside function.
			<var name1>, <var num2> := <var1 value>, <var2 value>
		NOTE: Whenever we define variable outside function we cannot use 3rd method
	*/
	w1, w2 := "hello", "World"
	num1, num2 := 5.6, 9.5
	var num3 int = 62

	fmt.Println(multiple(w1, w2))
	fmt.Println(add(num1, num2))

	// Changing datatype of variable
	var num4 float64 = float64(num3)
	fmt.Println(num4)
}
