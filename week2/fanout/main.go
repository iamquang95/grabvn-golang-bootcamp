package main

import (
	"fmt"
	"time"
)

func distributeInp(inp, mul1, mul2, mul3 chan int) {
	for value := range inp {
		mul1 <- value
		mul2 <- value
		mul3 <- value
	}
}

func handleOutput(inp chan int, f func(int) int) {
	for value := range inp {
		fmt.Println(f(value))
	}
}

func main() {
	inp := make(chan int)
	mul1 := make(chan int)
	mul2 := make(chan int)
	mul3 := make(chan int)

	go distributeInp(inp, mul1, mul2, mul3)
	go handleOutput(mul1, func(x int) int { return x * 1 })
	go handleOutput(mul2, func(x int) int { return x * 2 })
	go handleOutput(mul3, func(x int) int { return x * 3 })

	for i := 1; i <= 10; i++ {
		inp <- i
		time.Sleep(1 * time.Second)
	}

	fmt.Scanln()
}
