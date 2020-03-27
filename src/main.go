/*
Go is:
open source
by Google
expressive, concise, clean, efficient
inbuilt concurrency
modular approach
compiled language
garbage collected
static typed
fast!
*/

/*
Compile Demo
- cd /Users/nishak/Personal-Projects/Go/demo-compile
- go build hello.go
- ls -lha
- ./hello

See the size difference between hello and hello.class
> go build hello.go, size 71B
> creates a hello file, size 2.1M

While in java, 92B and 415B
*/

package main

import "fmt"

func main() {
	fmt.Println("=> Hello, 世界!")

	manyImports()

	randomNumber()
	exportedName()

	fmt.Println(swap(1, 2))
	fmt.Println(swap1(3, 4))

	variable()
	shortHandVariable()

	types()
	constants()
}
