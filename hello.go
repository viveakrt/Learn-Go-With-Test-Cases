package main

import "fmt"

const englishHelloPrefix = "Hello, "

func main() {
	fmt.Println(Hello("Vivek"))
}

func Hello(name string) string {
	return englishHelloPrefix + name
}
