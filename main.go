package main

import (
	app "delos/app"
	"fmt"
)

func main() {
	fmt.Println("==========Starting Server==========")
	app.Start()
	test()
}

func test() {
	var a = 5
	var p = &a // copy by reference
	var x = a  // copy by value

	fmt.Println("a = ", a)   // a =  5
	fmt.Println("&a = ", &a) // a =  5
	fmt.Println("p = ", p)   // p =  0x10414020
	fmt.Println("*p = ", *p) // *p =  5
	fmt.Println("&p = ", &p) // &p =  0x1040c128
	fmt.Println("x = ", x)   // x =  5
	fmt.Println("&x = ", &x) // &x =  0xc0001e0dc8

	fmt.Println("\n Change *p = 3")
	*p = 3
	fmt.Println("a = ", a)   // a =  3
	fmt.Println("p = ", p)   // p =  0x10414020
	fmt.Println("*p = ", *p) // *p =  3
	fmt.Println("&p = ", &p) // &p =  0x1040c128
	fmt.Println("x = ", x)   // x =  5

	fmt.Println("\n Change a = 888")
	a = 888
	fmt.Println("a = ", a)   // a =  888
	fmt.Println("p = ", p)   // p =  0x10414020
	fmt.Println("*p = ", *p) // *p =  888
	fmt.Println("&p = ", &p) // &p =  0x1040c128
	fmt.Println("x = ", x)   // x =  5

	fmt.Println("\n Change x = 1")
	x = 1
	fmt.Println("a = ", a)   // a =  888
	fmt.Println("p = ", p)   // p =  0x10414020
	fmt.Println("*p = ", *p) // *p =  888
	fmt.Println("&p = ", &p) // &p =  0x1040c128
	fmt.Println("x = ", x)   // x =  1

	// &p = 3 // error: Cannot assign to &p because this is the address of variable a
}
