// Command unitconv converts values between units of measurement.
package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/ASVATH2029/muck-target-unitconv/internal/units"
)

const usage = `unitconv - convert between units of measurement

usage:
  unitconv convert [-c CATEGORY] VALUE FROM TO
  unitconv list [CATEGORY]
  unitconv table CATEGORY UNIT
  unitconv categories

CATEGORY is inferred from FROM/TO when omitted, unless the two symbols
are ambiguous across categories.
`

func run(args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 || args[0] == "help" || args[0] == "-h" || args[0] == "--help" {
		fmt.Fprint(stdout, usage)
		return 0
	}
	switch args[0] {
	case "convert":
		return cmdConvert(args[1:], stdout, stderr)
	case "list":
		return cmdList(args[1:], stdout, stderr)
	case "table":
		return cmdTable(args[1:], stdout, stderr)
	case "categories":
		for _, c := range units.Categories() {
			fmt.Fprintln(stdout, c)
		}
		return 0
	default:
		fmt.Fprintf(stderr, "unknown command %q\n", args[0])
		fmt.Fprint(stderr, usage)
		return 2
	}
}

func cmdConvert(args []string, stdout, stderr io.Writer) int {
	fs := flag.NewFlagSet("convert", flag.ContinueOnError)
	fs.SetOutput(stderr)
	category := fs.String("c", "", "category (length, weight, temperature, data)")
	if err := fs.Parse(args); err != nil {
		return 2
	}
	rest := fs.Args()
	if len(rest) != 3 {
		fmt.Fprintln(stderr, "convert: expected VALUE FROM TO")
		return 2
	}
	value, err := strconv.ParseFloat(rest[0], 64)
	if err != nil {
		fmt.Fprintf(stderr, "convert: bad value %q\n", rest[0])
		return 2
	}
	from, to := rest[1], rest[2]

	cat := *category
	if cat == "" {
		found, err := inferCategory(from, to)
		if err != nil {
			fmt.Fprintf(stderr, "convert: %v (use -c to specify)\n", err)
			return 2
		}
		cat = found
	}

	result, err := units.Convert(cat, value, from, to)
	if err != nil {
		fmt.Fprintf(stderr, "convert: %v\n", err)
		return 1
	}
	fmt.Fprintf(stdout, "%s %s = %s %s\n", units.FormatValue(value), from, units.FormatValue(result), to)
	return 0
}

func inferCategory(from, to string) (string, error) {
	var matches []string
	for _, cat := range units.Categories() {
		u, err := units.Units(cat)
		if err != nil {
			continue
		}
		if containsFold(u, from) && containsFold(u, to) {
			matches = append(matches, cat)
		}
	}
	switch len(matches) {
	case 0:
		return "", fmt.Errorf("no category recognises both %q and %q", from, to)
	case 1:
		return matches[0], nil
	default:
		return "", fmt.Errorf("%q and %q are ambiguous across categories %s", from, to, strings.Join(matches, ", "))
	}
}

func containsFold(list []string, s string) bool {
	s = strings.ToLower(s)
	for _, x := range list {
		if x == s {
			return true
		}
	}
	return false
}

func cmdList(args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		for _, c := range units.Categories() {
			u, _ := units.Units(c)
			fmt.Fprintf(stdout, "%-12s %s\n", c, strings.Join(u, ", "))
		}
		return 0
	}
	u, err := units.Units(args[0])
	if err != nil {
		fmt.Fprintf(stderr, "list: %v\n", err)
		return 1
	}
	for _, x := range u {
		fmt.Fprintln(stdout, x)
	}
	return 0
}

func cmdTable(args []string, stdout, stderr io.Writer) int {
	if len(args) != 2 {
		fmt.Fprintln(stderr, "table: expected CATEGORY UNIT")
		return 2
	}
	category, unit := args[0], args[1]
	all, err := units.Units(category)
	if err != nil {
		fmt.Fprintf(stderr, "table: %v\n", err)
		return 1
	}
	if !containsFold(all, unit) {
		fmt.Fprintf(stderr, "table: unknown unit %q in category %q\n", unit, category)
		return 1
	}
	for _, target := range all {
		result, err := units.Convert(category, 1, unit, target)
		if err != nil {
			fmt.Fprintf(stderr, "table: %v\n", err)
			return 1
		}
		fmt.Fprintf(stdout, "1 %s = %s %s\n", unit, units.FormatValue(result), target)
	}
	return 0
}

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}
