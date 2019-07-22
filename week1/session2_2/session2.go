package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type expresion struct {
	A, B       float64
	Arthimetic string
}

func parse(text string) (expresion, error) {
	var err error
	expr := strings.Fields(text)
	if len(expr) != 3 {
		err = errors.New("Invalid input length")
		return expresion{}, err
	}
	a, err := strconv.ParseFloat(expr[0], 10)
	if err != nil {
		return expresion{}, err
	}
	b, err := strconv.ParseFloat(expr[2], 10)
	if err != nil {
		return expresion{}, err
	}
	arthimetic := expr[1]
	return expresion{a, b, arthimetic}, err
}

func add(exp expresion) float64 {
	return exp.A + exp.B
}

func sub(exp expresion) float64 {
	return exp.A - exp.B
}

func multiply(exp expresion) float64 {
	return exp.A * exp.B
}

func divide(exp expresion) (float64, error) {
	if exp.B == 0 {
		err := errors.New("Division by zero")
		return 0, err
	}
	return exp.A / exp.B, nil
}

func calculate(exp expresion) (res float64, err error) {
	switch exp.Arthimetic {
	case "+":
		res = add(exp)
	case "-":
		res = sub(exp)
	case "*":
		res = multiply(exp)
	case "/":
		res, err = divide(exp)
	default:
		err = errors.New("Invalid arthimetic")
	}
	return res, err
}

func logError(err error) {
	fmt.Println("Error: ", err)
}

func handleExpression(text string) (string, error) {
	exp, err := parse(text)
	if err != nil {
		return "", err
	}
	res, err := calculate(exp)
	resString := fmt.Sprintf("%v %s %v = %v\n", exp.A, exp.Arthimetic, exp.B, res)
	return resString, err
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("> ")
	for scanner.Scan() {
		text := scanner.Text()
		res, err := handleExpression(text)
		if err != nil {
			logError(err)
			fmt.Print("> ")
			continue
		}
		fmt.Printf(res)
		fmt.Print("> ")
	}
}
