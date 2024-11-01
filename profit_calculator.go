package main

func calculateProfit(revenue, expenses, taxRate float64) float64 {

	return (revenue - expenses) * (1 - taxRate/100)
}
