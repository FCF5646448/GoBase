package main

import (
	"strconv"
	"fmt"
)

func main() {
	str := "abcdefg"
	//an,err := strconv.Atoi(str)
	//if err != nil {
	//	fmt.Println("str is not a interger")
	//	return
	//}
	//fmt.Println(an)

	if _,err := strconv.Atoi(str); err != nil {
		fmt.Println("str is not a interger")
		//return
	}

	var v int = 50
	switch {
	case v > 49, v < 30 :
		fmt.Println("v > 49 or v < 30")
	default:
		fmt.Println("default")
	}

	i := 5
	for i > 0  {
		i = i - 1
		fmt.Println(i)
	}

	for in, val := range str {
		fmt.Printf("%c at %d in str\n",val,in)
	}

}
