package main

import (
	"fmt"
)

var a int
var f float64
var s string

func main() {
	//enter int
	fmt.Print("Enter int value: ")
	fmt.Scanln(&a)
	//enter string
	fmt.Print("Enter string value: ")
	fmt.Scanln(&s)
	//enter float64
	fmt.Print("Enter float64 value: ")
	fmt.Scanln(&f)
	//print int
	fmt.Println(a)
	//print string
	fmt.Println(s)
	//print float64
	fmt.Print(f)
}
