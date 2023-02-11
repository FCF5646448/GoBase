package main

import "fmt"

func main(){
	fmt.Println("数组、切片、字典")

	f0()

	f1()

	f2()

	f3()

	f4()
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

func f2() {
	var arr1 [6]int
	var slice1 []int = arr1[2:5] // item at index 5 not included!

	// load the array with integers: 0,1,2,3,4,5
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
	}

	// print the slice
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}

	fmt.Printf("The length of arr1 is %d\n", len(arr1))
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

	// grow the slice
	slice1 = slice1[0:4]
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

	// grow the slice beyond capacity
	//slice1 = slice1[0:7 ] // panic: runtime error: slice bound out of range
}

func f3() {
	arr := [4]int{1,2,3,4}
	var slice1 []int = arr[:1]
	var slice2 []int = arr[1:3]
	arr[0] = 0
	arr[1] = 5
	fmt.Println(slice1[0]) //0
	fmt.Println(slice2[0]) //5
	fmt.Println(arr) //0,5,3,4
	slice1[0] = 6
	slice2[0] = 7
	fmt.Println(arr) //6,7,3,4
}

func f4() {
	arr := [4]string{"a","b","c","d"}
	slice1 := arr[:3]
	slice2 := arr[3:]
	fmt.Println(cap(slice1)) //4
	fmt.Println(cap(slice2)) //1

	slice3 := arr[2:2]
	fmt.Println(len(slice3)) //0
}

