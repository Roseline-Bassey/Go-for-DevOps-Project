package main

import "fmt"

func main() {
// 	fmt.Println("1 + 1 =", 1 + 1)

// 	str := `this is a
// multiline
// string`
		
// 	fmt.Println(str)

// 	tabs := `this string
//             will have tabs 
//             in it`

// 	fmt.Println(tabs)
// 	fmt.Println(len("Hello world!"))
// 	fmt.Println("hello world"[2])
// 	fmt.Println("Hello " + "world")

// 	//boolean
// 	fmt.Println(true && true)
// 	fmt.Println(true && false)
// 	fmt.Println(true || true)
// 	fmt.Println(true || false)
// 	fmt.Println(!true)
	
	//Chapter 4 : Variables
	var x string = ("Hello World")
	fmt.Println(x)

	var a string = "Good morning"
	fmt.Println(a)

	var b string = ("first")
	// fmt.Println(b)
	b = ("second")
	fmt.Println(b)

	// Operators - Arithmetic Operators
	var c = 10 + 10
	var d = 10 - 20
	var e = 10 * 5
	var f = 10 / 2
	var g = 10 % 3
	// var h = c-- 
	// var i = c++
	fmt.Println(c,d,e,f,g)

	// Relational Operators
	var j = c == d
	var k = c != d
	var l = e > f
	var m = e < f
	var n = f >= g
	var o = f <= g
	fmt.Println(j,k,l,m,n,o)

}      
