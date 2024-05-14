package main

import "fmt"

/* global variable declaration */
var Car string = "Benz"

//MAIN FUNCTION

func main() {

	fmt.Println("1 + 1 =", 1 + 1)

	str := `this is a
multiline
string`
		
	fmt.Println(str)

	tabs := `this string
            will have tabs 
            in it`

	fmt.Println(tabs)
	fmt.Println(len("Hello world!"))
	fmt.Println("hello world"[2])
	fmt.Println("Hello " + "world")

	//boolean
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
	
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

	// // logical operators
	// var aa int = 10
	// var bb int = 20
	// var p = aa && bb
	// var q = bb || aa
	// // var r = !(d && c)
	// fmt.Println(p,q)

	z := 100
	fmt.Println(z)

	y := "Max"
	fmt.Println("My dog's name is", y)
    
	//global variable: The varible 'car' is declared outside of the main function.This can be access by other functions in the program
	fmt.Println(Car)
	
	//local variable: these are variables declared and initialized within a function or a block.
	var w, u, mm int

	w = 10
	u = 20
	mm = w + u

	fmt.Printf ("value of w = %d, u = %d and mm = %d\n", w, u, mm)

	// Formal parameters: These are treated as local variables within a function but they take preferance over local variables


   // Sample Program: This Go program takes user input, doubles it, and then prints the result. The user's input must be an int else it returns '0'
   fmt.Print("Tell us your name: ")
   var input float64
   fmt.Scanf("%f", &input)

   output := input * 2
   fmt.Println(output)

}      
