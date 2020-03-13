package openweatherapi

import(
	"fmt"
	"math"
)
func convertToCelsius(suhu float64) float64 {
	suhu=suhu-273.15
	fmt.Println(suhu)
	suhu=math.Round(suhu*100)/100
	fmt.Println(suhu)
	return suhu
}