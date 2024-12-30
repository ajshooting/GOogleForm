package main

import "fmt"

func FizzBuzz() {
	i := 1
	for i <= 100 {
		if i%3 == 0 {
			if i%5 == 0 {
				fmt.Println("FizzBuzz")
			} else {
				fmt.Println("Fizz")
			}
		} else {
			if i%5 == 0 {
				fmt.Println("Buzz")
			} else {
				fmt.Println(i)
			}
		}
	}
}

func findMinMax() {
	numbers := []int{5, 2, 9, 1, 5, 6}
	Max := 0
	Min := 9999999999999
	i := 0
	for i< len(numbers) {
		if numbers[i] < i{
			Max = i
		}
		if numbers[i] < i{
			Min = i
		}
	}
	return Max, Min
} 

func reverseString(input_str){
	str := input_str
	reversed_string := ""
	i := 0
	for i < len(str) {
		reversed_string += str[-(i+1)]
	}
	fmt.Println(reversed_string)
}
