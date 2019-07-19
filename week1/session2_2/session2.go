package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parse(text string) (a float64, b float64, arthimatic string, err error) {
	var expr = strings.Fields(text)
	if len(expr) != 3 {
		err = errors.New("invalid")
		return
	}
	a, err = strconv.ParseFloat(expr[0], 10)
	if err != nil {
		return
	}
	b, err = strconv.ParseFloat(expr[2], 10)
	arthimatic = expr[1]
	return
}

func calculate(a float64, b float64, arthimatic string) (res float64, err error) {
	switch arthimatic {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		if b == 0 {
			err = errors.New("division by zero")
			return
		} else {
			res = float64(a) / float64(b)
		}
	default:
		err = errors.New("Invalid arthimatic")
	}
	return
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("> ")
	for scanner.Scan() {
		text := scanner.Text()
		var a, b, arthimatic, err = parse(text)
		if err != nil {
			fmt.Println("Error")
		} else {
			res, err := calculate(a, b, arthimatic)
			if err != nil {
				fmt.Println("Error")
			} else {
				fmt.Printf("%v %s %v = %v\n", a, arthimatic, b, res)
			}
		}
		fmt.Print("> ")
	}
}
