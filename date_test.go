package vo

import "testing"

func TestFuzzyDate(t *testing.T) {
	valid := []string{
		"2024",
		"2024-03",
		"2024-03-15",
		"1999",
		"2000-01",
		"2000-12-31",
	}
	for _, v := range valid {
		if err := FuzzyDate("date", v); err != nil {
			t.Errorf("expected %q to be valid, got: %s", v, err)
		}
	}

	invalid := []string{
		"",
		"24",
		"2024-3",
		"2024-3-15",
		"2024-03-1",
		"2024/03/15",
		"03-2024",
		"not-a-date",
	}
	for _, v := range invalid {
		if err := FuzzyDate("date", v); err == nil {
			t.Errorf("expected %q to be invalid", v)
		}
	}
}
