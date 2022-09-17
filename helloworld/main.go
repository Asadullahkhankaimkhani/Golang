package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	var whattoSay string
	var i int

	whattoSay = "Hello, World!"
	fmt.Println(whattoSay)
	i = 1
	fmt.Println(i)
	unExported, unExpected2 := twoReturn()

	fmt.Println(unExported)
	fmt.Println(unExpected2)
}

func twoReturn() (int, int) {
	return 1, 2
}
