package main

import "fmt"

func main(){
	var s string = "This is string"

	var i int = 24

	var b bool = false

	var f float64 = 24.24
	
	//using Println :

	fmt.Println()
	fmt.Println(s, "\n")
	fmt.Println("var with integer: ", i, "\n")					// println + \n  = 2 line break
	fmt.Println("var with bool: ", b, "\n")
	fmt.Println("var with float64: ", f, "\n")

	//using Printf :

	fmt.Printf("Printing using printf method : %s\n", s)
														
	fmt.Printf("Printing multiple var using printf : %d ", i)

	println("\n")												//This adds 2 lines of line break.
	
	//using print :

	fmt.Print("print : ", s)
}


/*

- The line break or a new line can be used by "\n".

- fmt.Printf: It is used to print format specifiers.

- fmt.Println: It is used when a line break is needed after the output by default.

- fmt.Print: This gives the user the liberty to decide whether or not to include a line break after the output.

*/
