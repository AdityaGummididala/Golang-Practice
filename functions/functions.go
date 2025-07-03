package main. //every go program must start with a package declaration
               //main is a special package that tells go this is an executable program not an library

import (.   //import is used to bring in other packages 
	"fmt".   //fmt package provides formatting functions
)

func add(x,y int)int{.   //func declares the function, add is the function name
	return x+y           //(x,y int) means both are int value; int after the closing parenthesis means the return type 
}

func main(){           //this is the entry point of any go executable program, when you run go run execution starts from here
fmt.Println(add(4,5)).  //this is used to print text or values in the terminal

}

// Function is a block of reusable code that performs a specific task.

