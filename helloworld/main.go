package main

import "github.com/helloWorld/helpers"

var s = "Hello"

type User struct {
	Name string
}

type Animal interface {
	Speak() string
	NumberOfLegs() int
}

type Dog struct {
	Name string
}

type Gorilla struct {
	Name string
}

func main() {

	v1 := helpers.RandomNumber()

	println(v1)

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

	// var1 := myStruct{
	// 	Name: "John",
	// }

	// var var2 myStruct
	// var2.Name = "Wick"

	// log.Println(var1.demoFunction())
	// log.Println(var2.demoFunction())

	// myMap := make(map[string]User)

	// myMap["0"] = User{
	// 	Name: "John",
	// 	Age:  30,
	// }

	// log.Println(myMap["0"])

	// var mySlice []User

	// mySlice = append(mySlice, User{
	// 	Name: "John",
	// 	Age:  30,
	// })

	// log.Println(mySlice[0])\

	// var mySlice []int

	// mySlice = append(mySlice, 1)
	// mySlice = append(mySlice, 2)
	// mySlice = append(mySlice, 3)

	// sort.Ints(mySlice)

	// log.Println(mySlice)

	// names := []string{"John", "Wick", "Doe"}

	// for i, name := range names {
	// 	log.Println(i, name)
	// }

	// if true {
	// 	log.Println("true")
	// } else {
	// 	log.Println("false")
	// }

	// decision structure is same as C

	// switch statement

	// switch "John" {
	// case "John":
	// 	log.Println("John")
	// case "Wick":
	// 	log.Println("Wick")
	// default:
	// 	log.Println("Default")
	// }

	// for i := 0; i <= 10; i++ {
	// 	log.Println(i)
	// }

	// animals := []string{"Dog", "Cat", "Bird"}

	// for _, animal := range animals {
	// 	log.Println(animal)
	// }

	// dog := Dog{
	// 	Name: "Dog",
	// }

	// gorilla := Gorilla{
	// 	Name: "Gorilla",
	// }

	// PrintInfo(&dog)
	// PrintInfo(&gorilla)

	// v1, v2 := twoReturn()

	// log.Println(v1)
	// log.Println(v2)

}

// func PrintInfo(animal Animal) {
// 	fmt.Println(animal.Speak())
// 	fmt.Println(animal.NumberOfLegs())
// }

// func (dog *Dog) Speak() string {
// 	return "Woof"
// }

// func (dog *Dog) NumberOfLegs() int {
// 	return 4
// }

// func (gorilla *Gorilla) Speak() string {
// 	return "Ooh Ooh Ahhhelpers.RandomNumber

// func (gorilla *Gorilla) NumberOfLegs() int {
// 	return 2

// type myStruct struct {
// 	Name string
// }

// func (m *myStruct) demoFunction() string {
// 	return m.Name
// }

// func changeUsingPointer(myString *string) {
// 	*myString = "Red"
// }
