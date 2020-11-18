// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
)

var (
	expression = regexp.MustCompile(`\+|\-|\*|\/`)
	numbers    = regexp.MustCompile(`([0-9]*[.])?[0-9]+`)
)

// Expression calculates the values that are provided via string
func Expression(input string) (float64, error) {
	// exp, _ := regexp.MatchString("+/-*", input)
	exp := expression.FindString(input)
	// fmt.Println(exp)
	num := numbers.FindAllString(input, -1)
	firstNum, _ := strconv.ParseFloat(num[0], 64)
	secondNum, _ := strconv.ParseFloat(num[1], 64)
	var result float64
	switch exp {
	case "+":
		result = Add(firstNum, secondNum)
	case "-":
		result = Subtract(firstNum, secondNum)
	case "*":
		result = Multiply(firstNum, secondNum)
	case "/":
		result, err := Divide(firstNum, secondNum)
		return result, err
	}
	return result, nil
}

// Expression takes a formule in string form
// func Expression(input string) float64 {
// 	var expr string
// 	for _, v := range input {
// 		if v == "+" || v == "-" || v == "*" || v == "/" {
// 			expr = v
// 		}
// 	}

// 	var result float64

// 	switch true {
// 	case strings.Contains(input, "+"):
// 		result = Add(input[0], input[2])
// 	case strings.Contains(input, "-"):
// 		result = Subtract(input[0], input[2])
// 	case strings.Contains(input, "*"):
// 		result = Multiply(input[0], input[2])
// 	case strings.Contains(input, "/"):
// 		result, _ = Divide(input[0], input[2])
// 	}
// 	return result

// 	// input = strings.TrimSpace(input)
// 	// input = strings.Join(input, "")
// 	inputs := strings.Split(input, "")
// 	for i, v := range inputs {
// 		if v == " " {
// 			copy(inputs[i:], inputs[i+1:])
// 		}
// 	}
// 	fmt.Printf("%#v\n", inputs)
// 	// input = strings.Split(input, "")
// 	// inputs := strings .Fields(input)
// 	// for i, v := range inputs {
// 	// 	if v == " " {
// 	// 		copy(inputs[i:], inputs[i+1:])
// 	// 	}
// 	// 	if v == "\t" {
// 	// 		copy(inputs[i:], inputs[i+1:])
// 	// 	}
// 	// }

// 	// fmt.Println(inputs)
// 	// // inputs = strings.Split(inputs, "")
// 	// fmt.Println(inputs[0])

// 	var result float64
// 	first, _ := strconv.ParseFloat(inputs[0], 64)
// 	second, _ := strconv.ParseFloat(inputs[2], 64)
// 	switch exp := inputs[1]; exp {
// 	case "+":
// 		result = Add(first, second)
// 	case "-":
// 		result = Subtract(first, second)
// 	case "*":
// 		result = Multiply(first, second)
// 	case "/":
// 		result, _ = Divide(first, second)
// 	}
// 	return result
// }

// Add takes two numbers and returns the result of adding them together.
func Add(a, b float64) float64 {
	return a + b
}

// AddMany takes multiple numbers and returns the result of adding them together.
func AddMany(inputs ...float64) float64 {
	var result float64
	for _, input := range inputs {
		result += input
	}
	return result
}

// Subtract takes two numbers and returns the result of subtracting the second
// from the first.
func Subtract(a, b float64) float64 {
	return a - b
}

// SubtractMany takes multiple numbers and returns the result of subtracting them from eachother.
func SubtractMany(inputs ...float64) float64 {
	var result float64 = inputs[0]
	for i := 1; i < len(inputs); i++ {
		result -= inputs[i]
	}
	return result
}

// Multiply takes two numbers and returns the result of mulitplying the first
// number with the second number.
func Multiply(a, b float64) float64 {
	return a * b
}

// MultiplyMany takes two numbers and returns the result of mulitplying the first
// number with the second number.
func MultiplyMany(inputs ...float64) float64 {
	result := inputs[0]
	for i := 1; i < len(inputs); i++ {
		result *= inputs[i]
	}
	return result
}

// Divide takes two numbers and returns the result of dividing the first
// number with the second number.
func Divide(a, b float64) (result float64, returnError error) {
	if b == 0 {
		return 0, fmt.Errorf("bad input: %f %f (division by zero is undefined)", a, b)
	}
	result = a / b
	return result, nil
}

// DivideMany takes mulitple numbers and returns the result of dividing the
// number in order of the slice.
func DivideMany(inputs ...float64) (float64, error) {
	var result float64 = inputs[0]
	for i := 1; i < len(inputs); i++ {
		if inputs[i] == 0 {
			return 0, fmt.Errorf("bad input: %f (division by zero is undefined)", inputs[i])
		}
		result /= inputs[i]
	}
	return result, nil
}

// Sqrt takes one number and returns the result of Sqrt root of the number.
func Sqrt(a float64) (result float64, returnError error) {
	if a < 0 {
		return 0, fmt.Errorf("bad input: %f (Sqrt by negative number is not possible)", a)
	}
	result = math.Sqrt(a)
	return result, nil
}
