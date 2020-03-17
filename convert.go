package openweatherapi

import(
	"math"
)
func convertToCelsius(suhu float64) float64 {
	suhu=suhu-273.15
	suhu=math.Round(suhu*100)/100
	return suhu
}
