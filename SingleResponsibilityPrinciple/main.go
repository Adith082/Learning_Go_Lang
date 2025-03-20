package main

import "fmt"

func printHelloMessage() {
	fmt.Println("hello")
}

func getTwoNumbers() (int, int) {

	fmt.Println("give 2 numbers ")

	var num1 int
	var num2 int

	fmt.Scanln(&num1) // here & means "ampersand" which refers to address of particular variable
	fmt.Scanln(&num2)

	return num1, num2
}

func sumOfTwoNumbers(num1 int, num2 int) int {
	return (num1 + num2)
}

func printSum(sum int) {
	fmt.Println("result is ", sum)
}

func main() {
	printHelloMessage()
	num1, num2 := getTwoNumbers()
	sum := sumOfTwoNumbers(num1, num2)
	printSum(sum)
}
