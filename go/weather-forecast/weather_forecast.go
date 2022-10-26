// Package weather should make it easy to find the weather.
package weather

// CurrentCondition is a variable that stores the current condition.
var CurrentCondition string

// CurrentLocation is a variable that stores the current location.
var CurrentLocation string

// Forecast returns the weather condition for the city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
