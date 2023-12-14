package main

import "fmt"

func main(){

	//Declaration + Initialization
	var variable1 string
	variable1 = "HELLO,"

	//Same Datatype Variable:
	var variable2,variable3 string = "Namaste", "and Good Morning!!"

	//Printing the above declared variables:
	fmt.Printf("%s %s %s\n", variable1, variable2, variable3)


	//Multiple Datatype Variables:
	var(
		question string = "The square of 5 is :"
		answer int = 25
	)

	//Printing Multiple Variables:
	fmt.Printf("%s %d\n", question, answer)

	//Short method:
	name := "Radhe"
	fmt.Printf("Before: %s\n", name)
	name = "Krishna"	//Do not assign other Datatype 
	fmt.Printf("After: %s\n", name)
}