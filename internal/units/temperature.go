package units

import "fmt"

var temperatureUnits = map[string]bool{"c": true, "f": true, "k": true}

// TemperatureUnits returns the recognised temperature unit symbols.
func TemperatureUnits() []string { return []string{"c", "f", "k"} }

// ConvertTemperature converts value from one temperature unit to another.
// Temperature is affine, not purely multiplicative, so it goes through
// Celsius as a common base rather than using convertLinear.
func ConvertTemperature(value float64, from, to string) (float64, error) {
	if !temperatureUnits[from] {
		return 0, fmt.Errorf("unknown unit %q", from)
	}
	if !temperatureUnits[to] {
		return 0, fmt.Errorf("unknown unit %q", to)
	}
	return fromCelsius(toCelsius(value, from), to), nil
}

func toCelsius(v float64, from string) float64 {
	switch from {
	case "c":
		return v
	case "f":
		return FahrenheitToCelsius(v)
	case "k":
		return KelvinToCelsius(v)
	}
	return v
}

func fromCelsius(c float64, to string) float64 {
	switch to {
	case "c":
		return c
	case "f":
		return CelsiusToFahrenheit(c)
	case "k":
		return CelsiusToKelvin(c)
	}
	return c
}

// --- Named convenience wrappers ---------------------------------------
// CelsiusToFahrenheit, FahrenheitToCelsius, CelsiusToKelvin and
// KelvinToCelsius are reused above by the generic engine. The two direct
// Fahrenheit<->Kelvin shortcuts below are not: the engine always routes
// through Celsius, so these are unused convenience functions left over
// from writing "the complete set" of pairwise conversions.

func CelsiusToFahrenheit(c float64) float64 { return c*9/5 + 32 }
func FahrenheitToCelsius(f float64) float64 { return (f - 32) * 5 / 9 }
func CelsiusToKelvin(c float64) float64     { return c + 273.15 }
func KelvinToCelsius(k float64) float64     { return k - 273.15 }
func FahrenheitToKelvin(f float64) float64  { return CelsiusToKelvin(FahrenheitToCelsius(f)) }
func KelvinToFahrenheit(k float64) float64  { return CelsiusToFahrenheit(KelvinToCelsius(k)) }
