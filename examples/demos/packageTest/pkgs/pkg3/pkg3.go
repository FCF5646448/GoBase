package pkg3

import (
	"fmt"
)

var (
	_ = constInitCheck()
	_ = variableInit("v1")
	_ = variableInit("v2")
)

const (
	c1 = "c1"
	c2 = "c2"
)

func constInitCheck() string {
	if c1 != "" {
		fmt.Println("pkg3: const c1 has been initialized")
	}
	if c2 != "" {
		fmt.Println("pkg3: const c2 has been initialized")
	}
	return ""
}

func variableInit(name string) string {
	fmt.Printf("pkg3: var %s has been initialized\n", name)
	return name
}

func init() {
	fmt.Println("pkg3: first init func")
}

func init() {
	fmt.Println("pkg3: second init func")
}

func main() {
	fmt.Println("main func for pkg3 package")
}
