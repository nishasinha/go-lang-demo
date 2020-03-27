package main

import "fmt"

const Pi = 3.14

func constants() {
	fmt.Println("=> Inside constants")

	const local = true
	fmt.Println(Pi, local)
}
