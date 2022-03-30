package main

import "fmt"

func main() {
	//Insert Code here

	// activity 1
	// better to use var slice []float64 or slice := []float64{}
	slice := make([]float64, 4, 5)
	slice[0] = 9.50
	slice[1] = 8.00
	slice[2] = 10.20
	slice[3] = 7.40
	fmt.Println(len(slice), cap(slice))

	// activity 2
	fmt.Println(slice[2])
	slice[2] = 9.80
	fmt.Println(slice)

	// activity 3
	slice = append(slice, 8.40, 9.40, 7.20)
	fmt.Println(len(slice), cap(slice))

	// activity 4
	subSlice := slice[2:]
	fmt.Println(subSlice)
	fmt.Println(len(subSlice), cap(subSlice))

}
