package main

import "fmt"

func Hello(name string) string {
	return "Hello, " + name + "\n"
}

func main() {
	fmt.Printf(Hello("world"))
}
