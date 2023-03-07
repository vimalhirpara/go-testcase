package main

import "fmt"

func main() {
	result, _ := Multiply(2, 3)
	fmt.Println("Result:", result)
}

func Multiply(num1, num2 int) (int, error) {
	return num1 * num2, nil
}
