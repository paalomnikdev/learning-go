package main

import "fmt"

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	doubled := transformNumbers(&numbers, getTransformerFunction("double"))

	fmt.Println(doubled)

	fmt.Println(
		transformNumbers(&numbers, getTransformerFunction("triple")),
	)

	fmt.Println(
		transformNumbers(&numbers, getTransformerFunction("whatever")),
	)

	fmt.Println(
		transformNumbers(&numbers, createTransformer(5)),
	)

	fmt.Println("factorial: ", factorial(5))
	fmt.Println("sum: ", sumup(10, 20, 30))
	fmt.Println("sum from slice: ", sumup(55, numbers...))
}

func sumup(startingValue int, numbers ...int) int {
	sum := startingValue

	for _, v := range numbers {
		sum += v
	}

	return sum
}

func factorial(number int) int {
	if number == 0 {
		return 1
	}

	return number * factorial(number - 1)
}

func transformNumbers(numbers *[]int, transform transformFn) []int {
	tNumbers := []int{}

	for _, v := range *numbers {
		tNumbers = append(tNumbers, transform(v))
	}

	return tNumbers
}

func getTransformerFunction(transformationType string) transformFn {
	switch transformationType {
	case "double":
		return double
	case "triple":
		return triple
	default:
		return func(number int) int {
			return number
		}
	}
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}

func createTransformer(factor int) transformFn {
	return func(number int) int {
		return number * factor
	}
}
