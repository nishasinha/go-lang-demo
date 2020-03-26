package main

import "fmt"

// type removal at inputs
func swap(x, y int) (int, int) {
	fmt.Println("=> Inside swap function")
	return y, x
}

// Naked/ Named return
func swap1(x, y int) (a, b int) {
	fmt.Println("=> Inside swap 1 function")
	a = y
	b = x
	return
}
