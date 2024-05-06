package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int
	fmt.Scan(&a)
	x := a / 100
	y := a / 10 % 10
	z := a % 10
	fmt.Print(strconv.Itoa(z) + strconv.Itoa(y) + strconv.Itoa(x))
}
