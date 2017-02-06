package gg

import (
	"fmt"
)

//Calculator interface
type Calculator interface {
	Calculate(a int, b int) int
}

//Add struct
type Add struct {
}

//Calculate func return int
func (a Add) Calculate(x int, y int) int {
	return x + y
}

//Sub struct
type Sub struct {
}

//Calculate func return int
func (s Sub) Calculate(x int, y int) int {
	return x - y
}

func main() {
	var add Add
	var sub Sub

	var cal Calculator

	cal = add
	fmt.Println(cal.Calculate(1, 2))

	cal = sub
	fmt.Println(cal.Calculate(1, 2))
}
