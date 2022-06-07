package main

import "fmt"

type Employee struct {
	firstName string
	lastName  string
	age       int
	salary    int
	topSpeed  float64
}

func main() {

	mystruct := Employee{
		firstName: "John",
		lastName:  "Jorge",
		age:       21,
		salary:    120000,
		topSpeed:  18.2,
	}

	fmt.Printf("Name is: %s\n", mystruct.firstName)

	age := 21
	fmt.Printf("Age is: %d\n", age)

	fmt.Printf("Employee is %v\n", mystruct)
	fmt.Printf("Employee is %+v\n", mystruct)
	fmt.Printf("Employee is %#v\n", mystruct)

	fmt.Printf("Name: %s, Age: %d\n", mystruct.firstName, mystruct.age)
	fmt.Printf("Name: %s, Age: %d, Top Running Speed %2.1f\n", mystruct.firstName, mystruct.age, mystruct.topSpeed)
}
