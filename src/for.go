package main

import "fmt"

func for1() {
	fmt.Println("=> Inside for1()")
	sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Println("Sum is: ", sum)
}

func while() {
	fmt.Println("=> Inside while()")

	sum := 1
	for sum < 1000 {
		sum += sum
	}
}
