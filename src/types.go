/*
Types:
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
*/
package main

import (
	"fmt"
	"math/cmplx"
)

func types() {
	fmt.Println("=> Inside type()")
	var (
		aInt   int        = 1
		aFloat float32    = 1.2
		aBool  bool       = false
		aCmplx complex128 = cmplx.Sqrt(-5 + 12i)
	)

	fmt.Printf("Type: %T Value: %v\n", aInt, aInt)
	fmt.Printf("Type: %T Value: %v\n", aFloat, aFloat)
	fmt.Printf("Type: %T Value: %v\n", aBool, aBool)
	fmt.Printf("Type: %T Value: %v\n", aCmplx, aCmplx)

	// type conversion
	var x int32 = int32(aInt)
	fmt.Printf("Type: %T Value: %v\n", x, x)

	// auto type inference
	var y = (-5 + 2i)
	fmt.Printf("Type: %T Value: %v\n", y, y)
	z := 42
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
