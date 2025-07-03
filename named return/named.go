package main
/*
In go when you define a function, you can name the return values in the function signature.
This allows you to:
Refer to those variables inside the function body
Return them with just a return (no need to write variables again)
*/

import (
	"fmt"
)

func split(sum int) (x,y int){
	x = sum * 2
	y = sum * x
	return
}

func main(){
	fmt.Println(split(5))
}