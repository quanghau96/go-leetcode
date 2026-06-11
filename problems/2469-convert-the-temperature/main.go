package main

func convertTemperature(celsius float64) []float64 {
	Kelvin := celsius + 273.15
	Fahrenheit := celsius*1.80 + 32.00

	return []float64{Kelvin, Fahrenheit}
}

func main() {
	temp := convertTemperature(36.50)
	println(temp[0], temp[1])
}

// problem https://leetcode.com/problems/convert-the-temperature/
