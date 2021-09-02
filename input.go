package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/anshul47/bmi/config"
)

var reader = bufio.NewReader(os.Stdin)

func GetUserMetrics() (float64, float64) {
	weight := getUserFInput(config.WeightPrompt)
	height := getUserFInput(config.HeightPrompt)
	return weight, height
}

func getUserFInput(que string) float64 {
	fmt.Print(que)
	input, _ := reader.ReadString('\n')
	if runtime.GOOS == "windows" {
		input = strings.Replace(input, "\r\n", "", -1)
	} else {
		input = strings.Replace(input, "\n", "", -1)
	}
	enteredValue, _ := strconv.ParseFloat(input, 64)
	return enteredValue
}
