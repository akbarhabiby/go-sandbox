package main

import (
	"errors"
	"fmt"
)

func main() {
	// * Variable are initialized to their zero value
	fmt.Println("=== Declaring a variable and printing it ===")
	// var age int = 50
	age := 50
	age = 25
	fmt.Println("This is an age", age)

	fmt.Println("\n=== Conditional ===")
	if age < 18 {
		fmt.Println("If statement age under 18")
	} else if age < 20 {
		fmt.Println("If statement age under 20")
	} else {
		fmt.Println("If statement age above 20")
	}

	switch { // * Switch works without break and without condition (params), ex: switch(params) in JavaScript
	case age < 18:
		fmt.Println("Switch case age under 18")
	case age < 20:
		fmt.Println("Switch case age under 20")
	default:
		fmt.Println("Switch case age above 20")
	}

	// var command string = "hi"
	command := "hi" // * var command string = "hi"

	switch command { // * Switch with condition (params)
	case "hi":
		fmt.Println("Switch case & command", command)
	case "hello":
		fmt.Println("Switch case & command", command)
	default:
		fmt.Println("Switch case & command did not match", command)
	}

	fmt.Println("\n=== Functional ===")
	mixin := mixin("Akbar", "Coding")

	fmt.Println(mixin)

	greeting, err := greet("Akbar", "Coding", "Programmer")
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println(greeting)
	}

	fmt.Println("\n=== Declare and assign multiple things ===")
	javascript, python, golang := "node", "py", true
	fmt.Println(javascript, python, golang)

	fmt.Println("\n=== Looping ===")
	// * Go only have for loop, no while loop
	// for { // You have an infinite loop
	// 	fmt.Println("hello")
	// }

	for i := 0; i < 10; i++ { // for loop
		fmt.Println("for: Current value of i: ", i)
	}

	num := 0

	for num < 10 { // while loop
		fmt.Println("while: Current value of i: ", num)
		num++
	}

	// * Struct is a grouping of fields, like an Object in JavaScript
	type Person struct {
		name string
		age  int
	}

	akbar := Person{}
	// var akbar Person
	fmt.Printf("%+v", akbar)
	akbar.name = "Akbar"
	akbar.age = 19
	fmt.Printf("%+v", akbar)

	arif := Person{name: "Arif", age: 18}
	aji := Person{"Aji", 18}
	fmt.Printf("%+v", arif)
	fmt.Printf("%+v", aji)
}

func mixin(name string, hobby string) string {
	return name + "" + hobby
}

func greet(name, hobby, job string) (string, error) {
	if name != "" && hobby != "" && job != "" {
		return "Hello " + name + ". Your hobby is " + hobby + " and your job is " + job, nil
	} else {
		return "", errors.New("Name, hobby, or Job cannot be empty!")
	}
}

// ! OLD
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println(helloWorld())
// }

// func helloWorld() string {
// 	return "Hello world XD"
// }
