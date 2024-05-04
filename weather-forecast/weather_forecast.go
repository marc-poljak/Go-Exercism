// Package weather provides functionality to retrieve and display weather forecasts.
package weather

// CurrentCondition represents the current weather condition.
var CurrentCondition string

// CurrentLocation represents the current location for the weather forecast.
var CurrentLocation string

// Forecast returns a weather forecast for the specified city and condition.
// It updates the CurrentLocation and CurrentCondition variables with the provided city and condition.
// The forecast message includes the city and current weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
    return CurrentLocation + " - current weather condition: " + CurrentCondition
}
