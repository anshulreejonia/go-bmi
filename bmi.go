package main

import (
	"fmt"

	"github.com/anshul47/bmi/config"
)

func main() {
	config.PrintWelcome()
	weight, height := GetUserMetrics()
	bmi := getBMI(weight, height)
	fmt.Printf("Your bmi: %.2f", bmi)
}

func getBMI(weight float64, height float64) float64 {
	return (weight / (height * height))

}
