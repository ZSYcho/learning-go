package main

import (
	"fmt"
	"unicode"
)

func main() {

	//username := "test" // 4bytes ASCII string
	username := "表情包"
	fmt.Println(len(username))
	fmt.Println(username[0])
	fmt.Printf("%c\n", username[0])
	for _, v := range username {
		fmt.Println(string(v))
	}

	// the best way is use 'range' to loop output the string
	// and use 'rune'

	data := []rune{'表', '情', '包'}
	for _, v := range data {
		fmt.Println(string(v), unicode.IsLetter(v))
	}

}
