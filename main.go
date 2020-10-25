package main

import "fmt"

func add(number1 int, number2 int) int {
	return number1 + number2
}

func main() {
	number1, number2 := 5, 5;
	
	result := fmt.Sprintf("%d + %d = %d", number1, number2 , add(number1, number2))
	fmt.Println(result);
}