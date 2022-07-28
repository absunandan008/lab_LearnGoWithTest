package v1

import "testing"

func TestRomanNumerals(t *testing.T) {
	cases := []struct {
		Description string
		Arabic      int
		Want        string
	}{
		{"1 gets converted to I", 1, "I"},
		{"2 gets converted to II", 2, "II"},
	}

	for _, c := range cases {

		t.Run(c.Description, func(t *testing.T) {

			got := ConvertToRoman(c.Arabic)
			if got != c.Want {
				t.Errorf("got %q but want %q", got, c.Want)
			}

		})
	}

}

func ConvertToRoman(arabic int) string {
	if arabic == 2 {
		return "II"
	}
	return "I"
}
