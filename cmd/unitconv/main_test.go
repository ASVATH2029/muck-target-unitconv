package main

import (
	"bytes"
	"strings"
	"testing"
)

func exec(args ...string) (int, string, string) {
	var out, errb bytes.Buffer
	code := run(args, &out, &errb)
	return code, out.String(), errb.String()
}

func TestHelpAndCategories(t *testing.T) {
	code, out, _ := exec()
	if code != 0 || !strings.Contains(out, "usage:") {
		t.Fatalf("help: %d %q", code, out)
	}
	code, out, _ = exec("categories")
	if code != 0 || !strings.Contains(out, "length\n") {
		t.Fatalf("categories: %d %q", code, out)
	}
	code, _, errs := exec("bogus")
	if code != 2 || !strings.Contains(errs, "unknown command") {
		t.Fatalf("bogus: %d %q", code, errs)
	}
}

func TestConvertCommand(t *testing.T) {
	code, out, _ := exec("convert", "1", "km", "m")
	if code != 0 || !strings.Contains(out, "1000") {
		t.Fatalf("convert (inferred category): %d %q", code, out)
	}
	code, out, _ = exec("convert", "-c", "temperature", "0", "c", "f")
	if code != 0 || !strings.Contains(out, "32") {
		t.Fatalf("convert (explicit category): %d %q", code, out)
	}
	code, _, errs := exec("convert", "-c", "weight", "1", "kg", "bogus")
	if code != 1 || !strings.Contains(errs, "unknown unit") {
		t.Fatalf("convert bad unit: %d %q", code, errs)
	}
	code, _, errs = exec("convert", "abc", "km", "m")
	if code != 2 || !strings.Contains(errs, "bad value") {
		t.Fatalf("convert bad number: %d %q", code, errs)
	}
	code, _, _ = exec("convert", "1", "m")
	if code != 2 {
		t.Fatalf("convert arity: %d", code)
	}
	code, _, errs = exec("convert", "1", "bogus1", "bogus2")
	if code != 2 || !strings.Contains(errs, "no category recognises both") {
		t.Fatalf("convert unresolvable category: %d %q", code, errs)
	}
}

func TestListAndTable(t *testing.T) {
	code, out, _ := exec("list")
	if code != 0 || !strings.Contains(out, "length") {
		t.Fatalf("list: %d %q", code, out)
	}
	code, out, _ = exec("list", "weight")
	if code != 0 || !strings.Contains(out, "kg") {
		t.Fatalf("list weight: %d %q", code, out)
	}
	code, _, errs := exec("list", "bogus")
	if code != 1 {
		t.Fatalf("list bogus: %d %q", code, errs)
	}
	code, out, _ = exec("table", "length", "m")
	if code != 0 || !strings.Contains(out, "1 m =") {
		t.Fatalf("table: %d %q", code, out)
	}
	code, _, errs = exec("table", "length", "bogus")
	if code != 1 {
		t.Fatalf("table bad unit: %d %q", code, errs)
	}
	code, _, _ = exec("table", "length")
	if code != 2 {
		t.Fatalf("table arity: %d", code)
	}
}
