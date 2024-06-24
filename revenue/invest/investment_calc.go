package main

import (
	"fmt"
	"math"
)

func main() {
	const inflation = 0.01
	var investAmount int = 1000
	var interestRate float64 = 0.5
	var years float64 = 10
	var futureVal float64 = float64(investAmount) * math.Pow((1+interestRate), years)
	fmt.Println(futureVal)
	future_real := futureVal / (math.Pow(1+inflation, years))
	var profit float64 = futureVal - float64(investAmount)
	fmt.Sprintln("Profit is:", profit)
	fmt.Println(future_real, "is future real")
}
