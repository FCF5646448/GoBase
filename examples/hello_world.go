package main

import (
	"fmt"
	"runtime"
	"os"
	"math/rand"
	"time"
)

const c = 1 //常量
var v int = 5 //变量


type T struct {
	value float64
}

func init() {
	fmt.Println("init func")
}

func main() {

	//a1,b1 := 10,0
	//c := a1 / b1 //panic: runtime error: integer divide by zero
	//print(c)

	fmt.Println("hello world go")
	fmt.Println(add(c,v))

	t := new(T)
	t.value = 0.001
	t.Method()

	const ( //iota
		c0 = iota
		c1 = iota
		c2 			//如果c2的表达式语句和前一个一样，那就可以省略掉。
		c3
	)
	const c4 = iota

	fmt.Println(c0,c1,c2,c3,c4)//0,1,2,3,0

	//runtime包获取操作系统类型
	fmt.Println("current system is %s",runtime.GOOS)
	//通过OS包获取环境变量的值
	fmt.Println("gopath is %s", os.Getenv("PATH"))


	randomTest()
}

func add(a int, b int) int {
	return a + b
}

func (t T) Method() {
	fmt.Println(t.value)
}

// 随机数
func randomTest() {
	for i:=0;i<5 ; i++ {
		a := rand.Int() //生成一个整形随机数
		fmt.Printf("%d ",a)
	}
	fmt.Println()
	for i:=0;i<5 ; i++  {
		a := rand.Intn(8) //生成一个介于[0,8)的随机数
		fmt.Printf("%d ",a)
	}
	fmt.Println()
	times := int64(time.Now().Nanosecond()) //当前时间的纳秒级数字
	rand.Seed(times) //随机数种子
	for i:=0; i<5; i++  {
		fmt.Printf("%2.2f  ",100*rand.Float32())
	}
}

