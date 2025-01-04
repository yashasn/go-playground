package main

import "fmt"

func main() {

	// var dummy string = "dummy"
	// alternateDummy := "alternate dummy"

	// //You cannnot have unused variable in Go

	// fmt.Println("Hello world from", dummy)
	// fmt.Println("Hello world from " + alternateDummy)

	// if length := len(dummy); length > 2 {
	// 	fmt.Println("Simpler if block syntax")
	// }
	// fmt.Println("----------")
	// fmt.Printf("Result from Test function is %d .\n", TestFunction(5, 10))
	// fmt.Println("----------")
	// a, b := TestFunctionImplicitReturn(5, 10)

	// fmt.Printf("A = %d   B = %d .\n", a, b)
	// fmt.Println("----------")
	// structFunction()
	// fmt.Println("----------")
	// interfaceTest()
	// fmt.Println("----------")
	// ErrorTest()
	// fmt.Println("----------")
	// TestSlices()
	fmt.Println("----------")
	TestCostsByDay()
}

func TestFunction(a, b int) int {
	return a + b
}

func TestFunctionImplicitReturn(x, y int) (a, b int) {
	a = x
	b = y

	//naked return
	return
}
