package main

import "fmt"

// START OMIT
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)     //声明一个通道channel // HL
	go sum(s[:len(s)/2], c) // [7, 2, 8] // go关键字启动协程 // HL
	go sum(s[len(s)/2:], c) // [-9, 4, 0] // HL
	x, y := <-c, <-c        // 从channel接收数据 // HL
	fmt.Printf("x=%d, y=%d, x+y=%d", x, y, x+y)
}

// END OMIT
