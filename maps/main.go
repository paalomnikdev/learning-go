package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main()  {
	websites := map[string]string{
		"Google": "https://google.com",
		"AWS": "https://aws.com",
	}
	fmt.Println(websites)

	fmt.Println(websites["AWS"])
	websites["Linkedin"] = "https://linkedin.com"
	fmt.Println(websites)

	delete(websites, "Google")
	fmt.Println(websites)

	courseRatings := make(floatMap, 3)

	courseRatings["go"] = 2.7
	courseRatings["python"] = 5
	courseRatings["angular"] = 3

	courseRatings.output()

	for k, v := range courseRatings {
		fmt.Println(k)
		fmt.Println(v)
	}
}
