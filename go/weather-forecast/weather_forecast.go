// Package weather location and condition tell forecast.
package weather

// CurrentCondition weather condition.
var CurrentCondition string
// CurrentLocation weather location.
var CurrentLocation string

// Forecast returns kinda of magic.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
