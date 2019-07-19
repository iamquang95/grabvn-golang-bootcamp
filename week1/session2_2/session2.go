package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parse(text string) (a float64, b float64, arthimetic string, err error) {
	var expr = strings.Fields(text)
	if len(expr) != 3 {
		err = errors.New("Invalid input length")
		return
	}
	a, err = strconv.ParseFloat(expr[0], 10)
	if err != nil {
		return
	}
	b, err = strconv.ParseFloat(expr[2], 10)
	arthimetic = expr[1]
	return
}

func calculate(a float64, b float64, arthimetic string) (res float64, err error) {
	switch arthimetic {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		if b == 0 {
			err = errors.New("Division by zero")
			return
		} else {
			res = float64(a) / float64(b)
		}
	default:
		err = errors.New("Invalid arthimetic")
	}
	return
}

func logError(err error) {
	fmt.Println("Error: ", err)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("> ")
	for scanner.Scan() {
		text := scanner.Text()
		var a, b, arthimetic, err = parse(text)
		if err != nil {
			logError(err)
		} else {
			res, err := calculate(a, b, arthimetic)
			if err != nil {
				logError(err)
			} else {
				fmt.Printf("%v %s %v = %v\n", a, arthimetic, b, res)
			}
		}
		fmt.Print("> ")
	}
}
