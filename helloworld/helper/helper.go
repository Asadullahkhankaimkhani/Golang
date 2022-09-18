package helper

var s = "Hello"

type User struct {
	Name string
	Age  int
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

func twoReturn() (string, string) {

	return "Hello", "World"
}
