package main

import "fmt"

func convertFahrenheitToCelsius(F float32) float32 {
	return (F - 32) * (5.0 / 9.0)
}
func convertCelsiusToFahrenheit(C float32) float32 {
	return C*(9.0/5.0) + 32
}
func main() {

	fmt.Println(convertCelsiusToFahrenheit(29.5)) //85.1
	fmt.Println(convertCelsiusToFahrenheit(22.1)) //71.78
	fmt.Println(convertCelsiusToFahrenheit(12.3)) //54.14

	fmt.Println(convertFahrenheitToCelsius(convertCelsiusToFahrenheit(29.5))) //29.5                          //
	fmt.Println(convertFahrenheitToCelsius(convertCelsiusToFahrenheit(22.1))) //22.1
	fmt.Println(convertFahrenheitToCelsius(convertCelsiusToFahrenheit(12.3))) //12.3
}
