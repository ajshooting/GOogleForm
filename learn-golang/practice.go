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

func findMinMax(numbers []int) (int, int) { // スライスを引数として受け取る
	if len(numbers) == 0 {
		return 0, 0 // エラーハンドリング: 空のスライスの場合
	}
	min := numbers[0]             // スライスの最初の要素で初期化
	max := numbers[0]             // スライスの最初の要素で初期化
	for _, num := range numbers { // rangeを使ってスライスの各要素を反復処理
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}
	return min, max
}

func reverseString(input_str string) string {
	runes := []rune(input_str) // 文字列をruneのスライスに変換
	reversed_string := ""
	for i := len(runes) - 1; i >= 0; i-- { // インデックスを逆順に回す
		reversed_string += string(runes[i]) // runeを文字列に変換して追加
	}
	return reversed_string
}
