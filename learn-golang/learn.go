package main

import "fmt"

const pi = 3.14159265359

func main_learn() {
	fmt.Println("Hello, World!")
	var (
		a int = 1
		message string = "pi is"
	)
	b := 2
	fmt.Println(a + b)
	fmt.Println(message, pi)
	for i :=0; i<10; i++ {
		fmt.Println(i)
	}
	// そんな感じね
}