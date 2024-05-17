package main

import "fmt"

func main () {

// This Go program takes user input, doubles it, and then prints the result. The user's input must be a number else it returns '0'
fmt.Print("Enter a number: ")
var input float64
fmt.Scanf("%f", &input)

output := input * 2
fmt.Println(output)

// Go Program 2 //
// Get Fahrenheit temperatute from users

 fmt.Print("Enter Fahrenheit temperarture: ")
 var fahrenheit float64
 fmt.Scanf("%f", &fahrenheit)

 celsius:= (fahrenheit -32) * 5/9
 fmt.Printf("%.2f degrees Fahrenheit is equal to %.2f degrees Celsius\n", fahrenheit, celsius)



// CHAPTER 4 - TASKS BEGINS FROM HERE	

/* 1. What is the value of x after running: 
x := 5; x += 1? */

x := 5; x += 1
fmt.Printf("value of x = %d\n", x)
}



