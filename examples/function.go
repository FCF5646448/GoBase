package main

import (
	"strconv"
	"fmt"
	"io"
	"runtime"
	"log"
)

type Options struct {
	par1 int
	par2 string
	par3 float64
}

func trace(s string) string{
	fmt.Printf("entering: %s\n",s)
	return s
}
func untrace(s string){fmt.Printf("leaving: %s\n",s)}


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

	LABEL1:
	for i:=0; i<=5 ; i++  {
		for j:=0; j<=5 ; j++ {
			if j==4 {
				continue LABEL1
			}
			fmt.Printf("i = %d , j = %d\n",i,j)
		}
	}

	a := 1
	goto LABEL2
	fmt.Printf("a = %d",a)
LABEL2:
	b := 2
	b += a
	fmt.Printf("a = %d , b = %d\n",a,b)

	x1,x2 := getX1AndX2(1)
	fmt.Printf("x1 = %d , x2 = %d\n",x1,x2)


	geeting("hello","Joe","Anna","Eileen")

	p := Options{3,"1",4.0}
	f1(1,2,p)

	typecheck(a,1,1.0,"1",true)

	f2()

	f4()

	f5()

	f6(1,2)

	fmt.Println(f7(3,f8))

	f9()


	fmt.Println(f10())

	fmt.Println(f11(1)(2))

	f12()

}

func getX1AndX2(input int) (x1,x2 int) {
	fmt.Println("has call")
	return
}

func geeting(prefix string,who ...string) {
	fmt.Println(prefix)
	w := who
	for _,val := range w  {
		fmt.Println(val)
	}
}

func f1(a int,b int, p Options) {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(p.par1)
	fmt.Println(p.par2)
	fmt.Println(p.par3)
}

func typecheck(values ...interface{}) {
	for _, val := range values {
		switch v := val.(type) {
			case int:
				fmt.Println("int")
			case float64:
				fmt.Println("float64")
			case string:
				fmt.Println("tring")
			default:
				_ = v
				fmt.Println("default")
		}
	}
}

func f2(){
	fmt.Println("In f2 top")
	defer f3()
	fmt.Println("In f2 bottom")
}

func f3() {
	fmt.Println("in f3")
	i := 0
	defer fmt.Println(i)
	i++
	return
}

func f4(){
	for i:=0;i<5;i++{
		defer fmt.Println(i)
	}
}

func f5(){
	defer untrace(trace("f5"))
	fmt.Println(" on f5")
}

func f6(a int, b int) (c int, err error) {
	defer func() {
		fmt.Printf("f6(%d,%d) = %d , %v\n",a,b,c,err)
	}()
	return a+b,io.EOF
}


func f7(a int, f func(b int, c int) int) int {
	return  a + f(a, a+1)
}

func f8(a int,b int) int {
	return a + b
}

func f9(){
	sum := 0
	for i:=0;i<5 ;i++  {
		g := func(i int) {
			sum += i
		}
		g(i)
	}
	fmt.Printf("sum = %d\n",sum)
}

func f10()(ret int) {
	defer func(){
		ret++
	}()
	return 1
}

//
func f11(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}

func f12() {
	where := func() {
		_,file,line,_ := runtime.Caller(1)
		log.Printf("%s:%d",file,line)
	}
	fmt.Println("f12 top")
	where()
	fmt.Println("f12 middle")
	where()
	fmt.Println("f12 bottom")
}
