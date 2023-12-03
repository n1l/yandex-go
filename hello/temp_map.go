package main

import "fmt"

func main() {
	pricelist := map[string]int{
		"хлеб":     50,
		"молоко":   100,
		"масло":    200,
		"колбаса":  500,
		"соль":     20,
		"огурцы":   200,
		"сыр":      600,
		"ветчина":  700,
		"буженина": 900,
		"помидоры": 250,
		"рыба":     300,
		"хамон":    1500,
	}

	for title, price := range pricelist {
		if price > 500 {
			fmt.Println(title, price)
		}
	}

	fmt.Println()

	order := []string{"хлеб", "буженина", "сыр", "огурцы"}

	var order_cost int
	for _, title := range order {
		order_cost += pricelist[title]
	}

	fmt.Println(order_cost)
}
