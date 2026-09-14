package units

// dataFactors gives, for each unit, how many bytes equal one unit. Decimal
// (kb, mb, ...) and binary (kib, mib, ...) prefixes are both supported.
var dataFactors = map[string]float64{
	"b":   1,
	"kb":  1000,
	"mb":  1000 * 1000,
	"gb":  1000 * 1000 * 1000,
	"tb":  1000 * 1000 * 1000 * 1000,
	"kib": 1024,
	"mib": 1024 * 1024,
	"gib": 1024 * 1024 * 1024,
	"tib": 1024 * 1024 * 1024 * 1024,
}

// DataUnits returns the recognised data-size unit symbols, sorted.
func DataUnits() []string { return sortedKeys(dataFactors) }

// ConvertData converts value from one data-size unit to another.
func ConvertData(value float64, from, to string) (float64, error) {
	return convertLinear(dataFactors, value, from, to)
}

// --- Named convenience wrappers ---------------------------------------

func BytesToKB(b float64) float64   { return b / dataFactors["kb"] }
func KBToBytes(kb float64) float64  { return kb * dataFactors["kb"] }
func BytesToMB(b float64) float64   { return b / dataFactors["mb"] }
func MBToBytes(mb float64) float64  { return mb * dataFactors["mb"] }
func BytesToGB(b float64) float64   { return b / dataFactors["gb"] }
func GBToBytes(gb float64) float64  { return gb * dataFactors["gb"] }
func BytesToKiB(b float64) float64  { return b / dataFactors["kib"] }
func KiBToBytes(kib float64) float64 { return kib * dataFactors["kib"] }
func BytesToMiB(b float64) float64  { return b / dataFactors["mib"] }
func MiBToBytes(mib float64) float64 { return mib * dataFactors["mib"] }
func BytesToGiB(b float64) float64  { return b / dataFactors["gib"] }
func GiBToBytes(gib float64) float64 { return gib * dataFactors["gib"] }
