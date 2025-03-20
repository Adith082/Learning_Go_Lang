package main

import "fmt"

func add(a int, b int) {

	sum := a + b
	fmt.Println(sum)
}

func addWithReturn(a int, b int) int {

	sum := a + b
	return sum
}

func addWithReturn2(a int, b int) (int, int) {
	sum := a + b
	mul := a * b
	return sum, mul
}

func printMe(name string) string {
	return name
}

func main() {
	
	a := 2
	b := 34
	add(a, b)
	res := addWithReturn(a, b)
	fmt.Println(res)
	p, q := addWithReturn2(a, b)
	fmt.Println(p, "ds", q, printMe("adith"))
}

/*
 golang language data types can be divided
 into several types such as numeric
 (integer, float, complex), String and boolean (bool)
*/
