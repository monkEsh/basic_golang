// key value system
// more like JSON
package main

import "fmt"

func main() {
	fmt.Println("Maps in Go")
	// map[key_datatype]value_datatype : map just refrence type dont have values
	// to assign values we need make map : make(map[key_datatype]value_datatype)
	grades := make(map[string]float32)

	grades["Timmy"] = 42
	grades["Jess"] = 92
	grades["Sam"] = 67

	fmt.Println(grades)

	TimsGrade := grades["Timmy"]
	fmt.Println("Timmys grade :", TimsGrade)

	delete(grades, "Timmy")
	fmt.Println(grades)

	for name, grade := range grades {
		fmt.Println("Name :", name, ", Grade : ", grade)
	}
}
