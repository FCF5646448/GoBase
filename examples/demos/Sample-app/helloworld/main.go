package main

import (
	"fmt"
	"unicode/utf8"
)

type myInt int

// const n myInt = 10
// const m int = n + 1 // 编译报错

// var a myInt = 1
// var b int = a + 1 // 编译报错

const a = 1

var b myInt = 5
var c = a + b

func main() {
	str := "世界"
	fmt.Printf("the length of s = %d\n", len(str)) //
	for i := 0; i < len(str); i++ {
		fmt.Printf("0x%x ", str[i])
	}

	fmt.Println("the character count in s is", utf8.RuneCountInString(str))
	for _, ch := range str {
		fmt.Printf("0x%x ", ch)
	}

	fmt.Printf("a: %b; b: %b; c: %b", a, b, c)
}
