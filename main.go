package main

import "fmt"

func main() {
	fmt.Println(SayHello())
}

func SayHello() string {
	return "Hello World!"
}

func Greet(name string) string {
	return "Welcome " + name
}

func ReplyName(name string) string {
	return name
}
