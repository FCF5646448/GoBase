package main

import "fmt"

type Rect struct {
	width, length float64
}

func (rect *Rect) area() float64 {
	rect.width = 30
	rect.length = 20
	return  rect.width * rect.length
}

type Ground struct {
	rect Rect //结构体内部定义了一个矩形
	price float64
}

func (g Ground) totalPrice() float64 {
	return g.price * g.rect.width * g.rect.length
}

func swap(x,y *int) {
	*x,*y = *y, *x
}

type Phone interface {
	call()
	ring()
}

type Huawei struct {
	price float64
}

func (hw Huawei) call() {
	fmt.Println("huawei call")
}

func (hw Huawei) ring(){
	fmt.Println("huawei ring")
}

type IPhone struct {
	color string
}

func (ip IPhone) call() {
	fmt.Println("iphone call")
}
func (ip IPhone) ring(){
	fmt.Println("iphone ring")
}

func main() {
	//x := 10
	//y := 100
	//swap(&x,&y)
	//fmt.Println(x)
	//fmt.Println(y)

	//rect := Rect{width:64,length:100} //也可以直接按成员定义顺序赋值，比如Rect{64,100}
	//fmt.Println(rect.width * rect.length)
	//rect := Rect{20,10}
	//fmt.Println(rect.area())

	var phone Phone
	phone = new(Huawei)
	phone.call()
	phone.ring()

	phone = new(IPhone)
	phone.call()

	var phones = [4]Phone{
		Huawei{100},
		IPhone{"000000"},
		Huawei{200},
		IPhone{"f0f0f0"},
	}
	fmt.Println(phones)
}
