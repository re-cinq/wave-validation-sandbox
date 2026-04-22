package main

import "fmt"

func Greet(name string) string {
	return fmt.Sprintf("hello, %s", name)
}

// Add returns the sum of two integers.
func Add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(Greet("world"))
}
