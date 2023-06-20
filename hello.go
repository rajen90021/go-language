package main

import (
	"fmt"
	"strconv"
)

//package level
// var (
// 	a float32 = 12
// 	b string  = "adarsh"
// 	c string  = "aduu "
// )

func main() {
	//var a int //decelaration
	//a = 12    //intialization

	///declaration with initialization
	//var b float32 = 133
	//	c := 123
	//fmt.Println("c=", c)
	//	fmt.Printf("b=%v, c=%v", b, c)
	//fmt.Println(a, b, c)
	//fmt.Printf("%T,%v", a, a)
	fmt.Print(a)
	fmt.Print("-----------------")
	fmt.Print("-----------------")
	var a int = 32
	fmt.Printf("%T,%v", a, a)
	var b string
	b = strconv.Itoa(a)
	fmt.Printf(" %T,%v", b, b)

}
