package main

import "fmt"

func main() {
	fmt.Println(sumInt(1, 2, 3, 4))
}

func sumInt(n ...int) (int, int) {
	var sum, co int
	for _, elem := range n {
		sum = sum + elem
		co++
	}
	return co, sum
}

/*
func add(x, y int, firstName, lastName string) (int, string) {
	var z = x + y
	var fullName = firstName + " " + lastName
	return z, fullName
}
*/
