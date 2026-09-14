package units

// weightFactors gives, for each unit, how many kilograms equal one unit.
var weightFactors = map[string]float64{
	"kg":    1,
	"g":     0.001,
	"mg":    0.000001,
	"lb":    0.45359237,
	"oz":    0.0283495231,
	"stone": 6.35029318,
}

// WeightUnits returns the recognised weight unit symbols, sorted.
func WeightUnits() []string { return sortedKeys(weightFactors) }

// ConvertWeight converts value from one weight unit to another.
func ConvertWeight(value float64, from, to string) (float64, error) {
	return convertLinear(weightFactors, value, from, to)
}

// --- Named convenience wrappers ---------------------------------------

func KgToLbs(kg float64) float64      { return kg / weightFactors["lb"] }
func LbsToKg(lb float64) float64      { return lb * weightFactors["lb"] }
func KgToOunces(kg float64) float64   { return kg / weightFactors["oz"] }
func OuncesToKg(oz float64) float64   { return oz * weightFactors["oz"] }
func KgToStone(kg float64) float64    { return kg / weightFactors["stone"] }
func StoneToKg(st float64) float64    { return st * weightFactors["stone"] }
func GramsToKg(g float64) float64     { return g / 1000 }
func KgToGrams(kg float64) float64    { return kg * 1000 }
