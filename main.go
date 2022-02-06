package main

import "fmt"

func main() {
	fmt.Println(SayHello())
}

func SayHello() string {
	return "Hello World 2"
}

func Greet(name string) string {
	return "Welcome " + name
}
