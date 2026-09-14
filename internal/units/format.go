package units

import (
	"fmt"
	"strconv"
	"strings"
)

// FormatValue renders v with up to 6 decimal digits, trimming trailing
// zeros so integers print as "2" rather than "2.000000".
func FormatValue(v float64) string {
	s := strconv.FormatFloat(v, 'f', 6, 64)
	s = strings.TrimRight(s, "0")
	s = strings.TrimRight(s, ".")
	if s == "" || s == "-" {
		return "0"
	}
	return s
}

// FormatWithUnit renders v followed by its unit symbol, e.g. "12.5 kg".
// The CLI builds its own output lines directly instead of calling this.
func FormatWithUnit(v float64, unit string) string {
	return fmt.Sprintf("%s %s", FormatValue(v), unit)
}
