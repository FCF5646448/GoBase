package main

import "fmt"

func main(){
	fmt.Println("数组、切片、字典")

	f0()

	f1()
}

func f0() {
	a := [...]string{"a", "b", "c", "d"}
	for i := range a {
		fmt.Println("Array item", i, "is", a[i])
	}
}

func f1() {
	arr := new([5]int) //arr实际上是数组首地址指针
	arr[0] = 0
	arr[1] = 1
	arr[2] = 2
	arr[3] = 3
	arr[4] = 4

	var arr2 [5]int = *arr
	arr2[1] = 10
	fmt.Println(*arr)
	fmt.Println(arr2)
}