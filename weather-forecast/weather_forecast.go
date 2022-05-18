// Package weather contains two variables and one method.
package weather

// CurrentCondition stores current condition.
var CurrentCondition string

// CurrentLocation stores current location.
var CurrentLocation string

// Forecast method tries to forecast weather condition for the given city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
