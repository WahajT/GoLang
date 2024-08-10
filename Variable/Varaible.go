package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Hello World\n")

	var variable1 int      //declare int type variable "variable1"
	fmt.Println(variable1) //print value 0
	var variable2 string
	fmt.Println(variable2) //print ""
	variable1 = 123
	variable2 = "hello world"
	fmt.Println(variable1, variable2) //print 123 hello world
	var variable3 = 123.45
	fmt.Println(variable3) //print 123.45
	variable4 := false
	fmt.Println(variable4) //print false
	variable5, variable6 := "multiple", 2
	fmt.Println(variable5, variable6) //print multiple 2”
}
