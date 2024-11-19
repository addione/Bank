package practice

import (
	"math"
)

func calculateInvestment(amount int, years int, roi float64) (float64, float64) {
	const inflationRate = 6.9
	var futureValue float64
	roi = 5.5
	returnValue := float64(amount) * math.Pow(1+roi/100, float64(years))
	futureValue = returnValue / math.Pow(1+inflationRate/100, float64(years))
	return returnValue, futureValue
}
