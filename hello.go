package main

import "fmt"

const prefix = "Hello, "

func main() {
	fmt.Println(Hello())
}

func Hello() string {
	return prefix + "World!"
}

func HelloTo(name string) string {
	if name == "" {
		return prefix + "World!"
	}

	return prefix + name
}
