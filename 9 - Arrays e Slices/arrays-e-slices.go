package main

import (
	"fmt"
)

func main() {
	fmt.Println("Arrays e Slices")
	// ARRAYS
	var array1 [5]string
	array1[0] = "Kalu Dematto Gastaldelli,"
	array1[1] = "Rosi Dematto Gastaldelli,"
	array1[2] = "Nubia Dematto Gastaldelli,"
	array1[3] = "Roana Dematto Gastaldelli,"
	array1[4] = "Irene Dematto"
	fmt.Println(array1)

	array2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(array3)

	//SLICES
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice = append(slice, 18)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = 100
	fmt.Println(slice2)
}
