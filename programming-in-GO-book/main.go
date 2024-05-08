package main

import "fmt"

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

}      
