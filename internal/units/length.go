package units

// lengthFactors gives, for each unit, how many meters equal one unit.
var lengthFactors = map[string]float64{
	"m":  1,
	"km": 1000,
	"cm": 0.01,
	"mm": 0.001,
	"ft": 0.3048,
	"in": 0.0254,
	"yd": 0.9144,
	"mi": 1609.344,
}

// LengthUnits returns the recognised length unit symbols, sorted.
func LengthUnits() []string { return sortedKeys(lengthFactors) }

// ConvertLength converts value from one length unit to another.
func ConvertLength(value float64, from, to string) (float64, error) {
	return convertLinear(lengthFactors, value, from, to)
}

// --- Named convenience wrappers ---------------------------------------
// These exist for callers who want a direct function instead of going
// through ConvertLength with unit strings. The CLI itself only uses
// ConvertLength.

func MetersToFeet(m float64) float64        { return m / lengthFactors["ft"] }
func FeetToMeters(ft float64) float64       { return ft * lengthFactors["ft"] }
func MetersToInches(m float64) float64      { return m / lengthFactors["in"] }
func InchesToMeters(in float64) float64     { return in * lengthFactors["in"] }
func MetersToMiles(m float64) float64       { return m / lengthFactors["mi"] }
func MilesToMeters(mi float64) float64      { return mi * lengthFactors["mi"] }
func MetersToYards(m float64) float64       { return m / lengthFactors["yd"] }
func YardsToMeters(yd float64) float64      { return yd * lengthFactors["yd"] }
func MetersToKilometers(m float64) float64  { return m / lengthFactors["km"] }
func KilometersToMeters(km float64) float64 { return km * lengthFactors["km"] }
func MetersToCentimeters(m float64) float64 { return m / lengthFactors["cm"] }
func CentimetersToMeters(cm float64) float64 { return cm * lengthFactors["cm"] }
