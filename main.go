package main

import (
	"fmt"
)

const (
	sunday = iota
	monday
	tuesday
	wednesday
)

func main() {
	array := []int64{1, 2, 3, 4, 5}
	mySlice := make([]int64, len(array))
	copy(mySlice, array)
	fmt.Println(array)
	fmt.Println(mySlice)
	mySlice = append(mySlice, 333)
	mySlice = append(mySlice, 444)
	mySlice = append(mySlice, 555)
	mySlice = append(mySlice, 666)
	mySlice[0] = 5
	fmt.Println(array)
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
}
