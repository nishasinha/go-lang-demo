package main

import "fmt"

func branching() {
	fmt.Println("=> Inside branching()")
	i := 10

	if i := 11; i < 11 {
		fmt.Println(i, " is less than 11")
	} else {
		fmt.Println(i, "i is NOT less than 11")
	}

	fmt.Println(i)

}
