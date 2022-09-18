package main

import "log"

var s = "Hello"

type User struct {
	Name string
	Age  int
}

func main() {
	// fmt.Println("Hello, World!")

	// var whattoSay string
	// var i int

	// whattoSay = "Hello, World!"
	// fmt.Println(whattoSay)
	// i = 1
	// fmt.Println(i)
	// unExported, unExpected2 := twoReturn()

	// fmt.Println(unExported)
	// fmt.Println(unExpected2)

	// var myString string
	// myString = "Green"

	// fmt.Println(myString)

	// changeUsingPointer(&myString)

	// fmt.Println(myString)

	// var s2 = "six"

	//var s = "five"

	//log.Println(s2)
	//log.Println(s)

	// var user User
	// user.Name = "John"
	// user.Age = 30

	// log.Println(user)

	var1 := myStruct{
		Name: "John",
	}

	var var2 myStruct
	var2.Name = "Wick"

	log.Println(var1.demoFunction())
	log.Println(var2.demoFunction())

}

type myStruct struct {
	Name string
}

func (m *myStruct) demoFunction() string {
	return m.Name
}

// func changeUsingPointer(myString *string) {
// 	*myString = "Red"
// }

// func twoReturn() (int, int) {
// 	return 1, 2
// }
