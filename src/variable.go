package main

import "fmt"

var ch, str, int1 = 'a', "hello", 12

func variable() {
	var int2 int = 0
	fmt.Println("=> Inside variables")
	fmt.Println(ch, str, int1, int2)
}

func shortHandVariable() {
	fmt.Println("=> Inside short hand variables, only inside func")
	ch1, str1, int11 := 'a', "hello", 12
	fmt.Println(ch1, str1, int11)
}
