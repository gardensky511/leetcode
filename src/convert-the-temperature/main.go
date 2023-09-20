package convert_the_temperature

// topic 👉 https://leetcode.com/problems/convert-the-temperature/description/

func convertTemperature(celsius float64) []float64 {
	kelvin := celsius + 273.15
	fahrenheit := celsius*1.80 + 32.00

	return []float64{kelvin, fahrenheit}
}
