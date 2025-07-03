package main

import (
	"bhara/mathstuff"
	"fmt"
)

func main() {
	fmt.Println("Addition", mathstuff.Add(2, 3)) //this is external name and this would run
	fmt.Println("subtraction", mathstuff.sub(2, 3)) //this is internalname and this would throw an error
}
