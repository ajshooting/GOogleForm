package main

import "fmt"

// スライスってsliceで合ってんのかな...

func slice_practice() {
	// これは配列(固定長)
	arr := [5]int{1, 2, 3, 4, 5}

	// これはスライス(可変長)
	numbers := []int{5, 2, 9, 1, 5, 6}

	// 配列からスライスを作成..?
	s1 := arr[1:4]

	fmt.Println(arr, numbers, s1)
}
