// Package weather provides tools that can forecast the current weather condition of various cities in Goblinocus.
package weather


var (
    // CurrentCondition shows the country's current weather condition.
	CurrentCondition string
    // CurrentLocation shows the country's geographical location.
	CurrentLocation  string
)

// Forecast shows the current weather forecast for the provided location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
