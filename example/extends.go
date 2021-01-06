package main

import (
	"fmt"
)

// START OMIT
type A struct {
}

func (*A) Hello(name string) {
	fmt.Println("hello " + name + ", i am a")
}

type B struct {
	*A
}

func (*B) Hello(name string) {
	fmt.Println("hello " + name + ", i am b")
}
func main() {
	name := "Lee"
	a := A{}
	a.Hello(name)
	b := B{&A{}}
	b.Hello(name)
}

// END OMIT
