package main

import (
	"fmt"

	"Goex2/myutils"
)

func main() {
	text := "Hello World"
	number := 5
	base := 2
	exponent := 3

	fmt.Println("Original String:", text)
	fmt.Println("Reversed String:", myutils.Reverse(text))
	fmt.Println("Number of Vowels:", myutils.CountVowels(text))

	fmt.Println("Factorial of", number, ":", myutils.Factorial(number))
	fmt.Println(base, "raised to the power", exponent, ":", myutils.Power(base, exponent))
}