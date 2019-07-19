package main

import "fmt"

func calculate() {
	var a, b int
	var arthimatic int32
	fmt.Printf("> ")
	fmt.Scanf("%d %c %d", &a, &arthimatic, &b)
	var res interface{}
	switch arthimatic {
	case '+':
		res = a + b
	case '-':
		res = a - b
	case '*':
		res = a * b
	case '/':
		res = float64(a) / float64(b)
	}
	fmt.Printf("%v %c %v = %v\n", a, arthimatic, b, res)
}

func main() {
	for {
		calculate()
	}
}
