package units

import "testing"

func TestConvertLength(t *testing.T) {
	got, err := ConvertLength(1, "km", "m")
	if err != nil || got != 1000 {
		t.Fatalf("got %v %v", got, err)
	}
	if _, err := ConvertLength(1, "km", "bogus"); err == nil {
		t.Fatal("expected error for unknown target unit")
	}
	if _, err := ConvertLength(1, "bogus", "km"); err == nil {
		t.Fatal("expected error for unknown source unit")
	}
}

func TestConvertWeight(t *testing.T) {
	got, err := ConvertWeight(1, "kg", "g")
	if err != nil || got != 1000 {
		t.Fatalf("got %v %v", got, err)
	}
}

func TestConvertData(t *testing.T) {
	got, err := ConvertData(1, "kb", "b")
	if err != nil || got != 1000 {
		t.Fatalf("decimal: got %v %v", got, err)
	}
	got, err = ConvertData(1, "kib", "b")
	if err != nil || got != 1024 {
		t.Fatalf("binary: got %v %v", got, err)
	}
}

func TestConvertTemperature(t *testing.T) {
	cases := []struct {
		v        float64
		from, to string
		want     float64
	}{
		{0, "c", "f", 32},
		{100, "c", "f", 212},
		{32, "f", "c", 0},
		{0, "c", "k", 273.15},
		{273.15, "k", "c", 0},
	}
	for _, c := range cases {
		got, err := ConvertTemperature(c.v, c.from, c.to)
		if err != nil {
			t.Fatal(err)
		}
		if diff := got - c.want; diff > 1e-9 || diff < -1e-9 {
			t.Errorf("%v %s->%s = %v, want %v", c.v, c.from, c.to, got, c.want)
		}
	}
	if _, err := ConvertTemperature(0, "x", "c"); err == nil {
		t.Fatal("expected error for unknown source unit")
	}
	if _, err := ConvertTemperature(0, "c", "x"); err == nil {
		t.Fatal("expected error for unknown target unit")
	}
}

func TestCategoriesAndUnits(t *testing.T) {
	if got := Categories(); len(got) != 4 {
		t.Fatalf("Categories() = %v", got)
	}
	if _, err := Units("bogus"); err == nil {
		t.Fatal("expected error")
	}
	u, err := Units(Length)
	if err != nil || len(u) == 0 {
		t.Fatalf("Units(length) = %v %v", u, err)
	}
}

func TestConvertDispatch(t *testing.T) {
	if _, err := Convert("bogus", 1, "m", "ft"); err == nil {
		t.Fatal("expected error for unknown category")
	}
	got, err := Convert(Weight, 1, "kg", "g")
	if err != nil || got != 1000 {
		t.Fatalf("got %v %v", got, err)
	}
}

func TestFormatValue(t *testing.T) {
	cases := map[float64]string{1.5: "1.5", 2.0: "2", 0: "0", -0.25: "-0.25"}
	for in, want := range cases {
		if got := FormatValue(in); got != want {
			t.Errorf("FormatValue(%v) = %q, want %q", in, got, want)
		}
	}
}
