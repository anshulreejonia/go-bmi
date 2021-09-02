package config

import "fmt"

const MainTitle = "BMI Calculator"
const Separator = "_________________"
const WeightPrompt = "Please enter you weight (kg): "
const HeightPrompt = "Please enter you height (m): "

func PrintWelcome() {
	fmt.Println(MainTitle)
	fmt.Println(Separator)
}
