package main

import "fmt"

func main(){

	//Array:
	arr := [] int{5,6,7,8,9}
	fmt.Println("Initial array:", arr)
	
	//Slicing:
	slice := arr[1:3]
	fmt.Println("After slicing:", slice)

	//Features of slice:
	fmt.Println("Slice Capacity:", cap(slice))	// capacity = (capacity of underlying array - starting index of the slice)
	fmt.Println("Slice Length:", len(slice))

	/*
	Slice is the supart of an array that is made defining limits.
	A slice does not owns any memory, it only has a reference to the original data.
	Hence, if any change is done to slice elements then the underlyning array will be affected.
	
	In other words slice acts as pointer that holds memory address for the original array for a defined range.
	*/

	//New array:
	newarr := []int{1,2,3,4,5,6,7,8,9}
	fmt.Println("\nNew Array:", newarr)
	newslice := newarr[1:5]
	fmt.Println("2nd Slice:", newslice)
	fmt.Println("newslice-len:", len(newslice))	
	fmt.Println("newslice-cap:", cap(newslice))

	//Appending to new slice:
	sliceappend := append(newslice, 90)
	fmt.Println("\nAfter append:", sliceappend)
	fmt.Println("sliceappend-len:", len(sliceappend))
	fmt.Println("sliceappend-cap:", cap(sliceappend))

	// currently in newslice the capacity is 8 and length 5.
	// now we'll extend the capacity limit.

	sliceappend2 := append(newslice, 80, 70, 60, 40, 30)	// number of append elements more than capacity.
	fmt.Println("\nAfter excess append:", sliceappend2)
	fmt.Println("sliceappend2-len:", len(sliceappend2))	
	fmt.Println("sliceappend2-cap:", cap(sliceappend2))		// Here the capactiy gets doubled to the initial capacity of slice.


	//Integrating multiple slice:
	new_slice := append(sliceappend, sliceappend2...)
	fmt.Println("\nIntegrating 2 slice:", new_slice)
}