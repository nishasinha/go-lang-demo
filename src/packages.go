/*
Every Go program has to be within a package.package tour

The main method inside the main package is the point where execution starts."
*/

package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func randomNumber() {
	fmt.Println("=> Inside randomNumber()")

	rand.Seed(time.Now().Unix())
	fmt.Println("My favorite number is", rand.Intn(30))
	fmt.Println("My favorite number is", rand.Intn(30))
	fmt.Println("My favorite number is", rand.Intn(30))
}

func exportedName() {
	fmt.Println("=> Inside exported names")
	// exported names start with Capital letter
	fmt.Println("Pi is", math.Pi)
}
