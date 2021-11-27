package main

import "fmt"

func main() {

	// Creating slices
	slc1 := []int{58, 69, 40, 45, 11, 56, 67, 21, 65}
	var slc2 []int
	slc3 := make([]int, 5)
	slc4 := []int{78, 50, 67, 77}

	// Before copying
	fmt.Println("Slice_1:", slc1)
	fmt.Println("Slice_2:", slc2)
	fmt.Println("Slice_3:", slc3)
	fmt.Println("Slice_4:", slc4)

	// Copying the slices
	copy_1 := copy(slc2, slc1)
	fmt.Println("\nSlice:", slc2)
	fmt.Println("Total number of elements copied:", copy_1)

	copy_2 := copy(slc3, slc1)
	fmt.Println("\nSlice:", slc3)
	fmt.Println("Total number of elements copied:", copy_2)

	copy_3 := copy(slc4, slc1)
	fmt.Println("\nSlice:", slc4)
	fmt.Println("Total number of elements copied:", copy_3)

	// Don't confuse here, because in above
	// line of code the slc4 has been copied
	// and hence modified permanently i.e.
	// slc 4 contains [58 69 40 45]
	copy_4 := copy(slc1, slc4)
	fmt.Println("\nSlice:", slc1)
	fmt.Println("Total number of elements copied:", copy_4)

}
