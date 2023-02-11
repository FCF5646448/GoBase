package main

//var a = "G"
//
//func main()  {
//	n()
//	m()
//	n()
//}
//
//func n() {print(a)}
//
//func m() {
//	a := "o"
//	print(a)
//}

//var a = "G"
//
//func main()  {
//	n()
//	m()
//	n()
//}
//
//func n() {
//	print(a)
//}
//
//func m() {
//	a = "o"
//	print(a)
//}

var a string

func main() {
	a = "G"
	print(a)
	f1()
}

func f1() {
	a := "O"
	print(a)
	f2()
}

func f2() {
	print(a)
}