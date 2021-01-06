package main

import (
	"fmt"
	"unsafe"
)

func main() {
	part := struct {
		a bool
		d int64
		c int8
		b int32
		e byte
	}{}
	fmt.Printf("part size: %d, align: %d\n", unsafe.Sizeof(part), unsafe.Alignof(part))
}
