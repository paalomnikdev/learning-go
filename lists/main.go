package main

import "fmt"

type Product struct {
	id int
	title string
	price float64
}

func main()  {
    hobbies := [3]string{"cars", "drones", "wealth"}
	fmt.Println(hobbies)
	fmt.Println("First element value: ", hobbies[0])
	fmt.Println("Second and 3rd values: ", hobbies[1:])
	primaryHobbies := hobbies[:2]
	fmt.Println("First and 2nd values: ", primaryHobbies)
	primaryHobbies = hobbies[1:3]
	fmt.Println("First and last hobbies: ", primaryHobbies)

	goals := []string{"learn", "grow", "evolve"}
	fmt.Println(goals)

	car := Product{
		1,
		"Nissan Rogue",
		99.99,
	}
	book := Product{
		2,
		"Learning Go",
		1.99,
	}

	products := []Product{car, book}
	fmt.Println(products)
	products = append(products, Product{
		3,
		"iPhone 16 ProMax",
		1000,
	})
	fmt.Println(products)
}

func mainSlices()  {
	prices := []float64{10.99, 8.99}
	fmt.Println(prices[0])
	fmt.Println(prices[1:])

	prices = append(prices, 5.99)

	fmt.Println(prices)
}

func mainInitial() {
	productNames := [4]string{"book one"}
	prices := [4]float64{10.9, 9.99, 45.99, 20.0}
	fmt.Println(prices)

	productNames[2] = "carpet"

	fmt.Println(productNames)

	fmt.Println(prices[2])

	featuredPrices := prices[1:3]
	fmt.Println(featuredPrices)
	discountPrices := prices[:3]
	fmt.Println(discountPrices)
	highlightedPrices := discountPrices[:1]
	fmt.Println(highlightedPrices)
	fmt.Println(len(highlightedPrices), cap(highlightedPrices))

	highlightedPrices = highlightedPrices[:3]
	fmt.Println(highlightedPrices)
	fmt.Println(len(highlightedPrices), cap(highlightedPrices))
}
