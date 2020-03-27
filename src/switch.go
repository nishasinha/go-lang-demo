/*
- no BREAK is required in each case
- case need not to be constants
- case value need not to be integers
- unlike scala, all case values should have one type
*/

package main

import "fmt"

func switchBranch(name string) {
	fmt.Println("Inside switch()", name)

	switch defaultName := "Alex"; name {
	case defaultName:
		{
			fmt.Println("default", defaultName)
		}
	case "Alex":
		{
			fmt.Println("Name is" + name)
		}
	case "Bob":
		{
			fmt.Println(name)
		}
	default:
		{
			fmt.Println(name)
		}

	}
}
