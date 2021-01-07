package main

import (
	"fmt"
	"unsafe"
)

func main() {
	part1 := struct {
		a bool
		d int64
		c int8
		b int32
		e byte
	}{}
	fmt.Printf("part1 size: %d, align: %d\n", unsafe.Sizeof(part1), unsafe.Alignof(part1))
	part2 := struct {
		a bool
		e byte
		c int8
		b int32
		d int64
	}{}
	fmt.Printf("part1 size: %d, align: %d\n", unsafe.Sizeof(part2), unsafe.Alignof(part2))
}
