package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Pi)

}

/*
In go lang, an exported name starts with upper case letter whereas an unexported name starts with lower case
Exported means This name is available to other packages (public)
Imported means you're bringing in another package using import.
When you import a package you can only access its exported names 
Please refer the module and example program
*/