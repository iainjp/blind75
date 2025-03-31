package exercises

func maxProfit(prices []int) int {
	minPrice := 100000000
	profit := 0

	for _, price := range prices {
		if price < minPrice {
			minPrice = price
			continue
		}

		if price-minPrice > profit {
			profit = price - minPrice
		}
	}

	return profit
}
