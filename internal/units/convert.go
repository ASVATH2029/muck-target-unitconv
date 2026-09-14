// Package units converts values between units of measurement across four
// categories: length, weight, temperature and data size.
package units

import (
	"fmt"
	"sort"
	"strings"
)

// Category names recognised by Convert and Units.
const (
	Length      = "length"
	Weight      = "weight"
	Temperature = "temperature"
	Data        = "data"
)

// Categories lists the recognised category names, sorted.
func Categories() []string {
	return []string{Data, Length, Temperature, Weight}
}

// Units returns the unit symbols recognised for a category, sorted.
func Units(category string) ([]string, error) {
	switch strings.ToLower(category) {
	case Length:
		return LengthUnits(), nil
	case Weight:
		return WeightUnits(), nil
	case Temperature:
		return TemperatureUnits(), nil
	case Data:
		return DataUnits(), nil
	default:
		return nil, fmt.Errorf("unknown category %q", category)
	}
}

// Convert converts value from one unit to another within category.
func Convert(category string, value float64, from, to string) (float64, error) {
	switch strings.ToLower(category) {
	case Length:
		return ConvertLength(value, from, to)
	case Weight:
		return ConvertWeight(value, from, to)
	case Temperature:
		return ConvertTemperature(value, from, to)
	case Data:
		return ConvertData(value, from, to)
	default:
		return 0, fmt.Errorf("unknown category %q", category)
	}
}

func sortedKeys(m map[string]float64) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

// convertLinear converts between units whose relationship to a common base
// unit is a simple multiplicative factor (everything except temperature).
func convertLinear(factors map[string]float64, value float64, from, to string) (float64, error) {
	ff, ok := factors[strings.ToLower(from)]
	if !ok {
		return 0, fmt.Errorf("unknown unit %q", from)
	}
	tf, ok := factors[strings.ToLower(to)]
	if !ok {
		return 0, fmt.Errorf("unknown unit %q", to)
	}
	return value * ff / tf, nil
}
