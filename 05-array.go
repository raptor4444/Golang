package main

import "fmt"

func main(){
	
	//array with fixed size :
	var arr1 [5] int = [5] int {1,2,3,4,5};
	arr2 := [5] int {1,2,3,4,5}

	//array with dynamic size :	
	//The compiler will implicitly will calculate length based on the number of elemets specified.
	arr3 := [...] int{1,2,3,4,5,6,7,8,9,10}

	//Printing arrays :
	fmt.Println("\narr1 :", arr1)
	fmt.Println("arr2 :", arr2)
	fmt.Println("arr3 :", arr3)

	//Printing length of arrays :
	fmt.Println("\nLength (arr1) :", len(arr1))
	fmt.Println("Length (arr2) :", len(arr2))
	fmt.Println("Length (arr3) :", len(arr3))

	//Printing the target values using index :
	fmt.Println("\nValue at index 2 in arr1 :", arr1[2]);

	//Changing elements using index :
	arr1[2] = 100;
	fmt.Println("\nAfter changing element value at index 2 to 100 :", arr1[2]);

	fmt.Println("\n");

	//Looping through array :
	for i:=0; i<len(arr1); i++{
		fmt.Println(arr1[i]);
	}
}