// All the programs in go lang are made up of packages
package main //programs start running in package main

import (
	"fmt"
	"math"
	"math/rand"
)

// Here we have imported fmt and math/rand packages; fmt provides us with formatted input output operations where as math/rand provides pseduo random number generators
func main() { //pseudo random number is a number that isn't random but it is random enough for most purposes
	fmt.Println("This is a number", rand.Intn(5))
	fmt.Printf("Now you have %g problems. \n", math.Sqrt(4))
}

// this is a single linecomment

/*
This is a mutliline comment
*/
