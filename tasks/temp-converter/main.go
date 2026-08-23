package main

import "fmt"

func getTempInput() float64 {
	var fTemp float64

	fmt.Println("Enter temprature in Fahrenheit:")
	_, err := fmt.Scan(&fTemp)

	if err != nil {
		fmt.Println("Invalid input provided.")
		return getTempInput()
	}

	return fTemp
}

func convertTemp(fTemp float64) float64 {
	cTemp := (fTemp - 32) * (5.0 / 9.0)

	return cTemp
}

func main() {
	input := getTempInput()

	temp := convertTemp(input)

	fmt.Printf("%.2f°F = %.2f°C\n", input, temp)

}
