package main

import "testing"

func TestConvertToRomanNumber(t *testing.T) {
	numbersToTest := []struct {
		arabicNumber int
		romanNumber  string
	}{
		{1, "I"},
		{2, "II"},
		{3, "III"},
		{4, "IV"},
		{5, "V"},
		{6, "VI"},
		{7, "VII"},
		{8, "VIII"},
		{9, "IX"},
		{10, "X"},
		{11, "XI"},
		{14, "XIV"},
		{15, "XV"},
		{16, "XVI"},
		{19, "XIX"},
		{20, "XX"},
		{30, "XXX"},
		{40, "XL"},
		{50, "L"},
		{60, "LX"},
		{70, "LXX"},
		{80, "LXXX"},
		{90, "XC"},
		{100, "C"},
		{200, "CC"},
		{300, "CCC"},
		{400, "CD"},
		{500, "D"},
		{600, "DC"},
		{900, "CM"},
	}

	for _, tt := range numbersToTest {
		assertConvertToRomanNumber(t, tt.arabicNumber, tt.romanNumber)
	}

}

func assertConvertToRomanNumber(t *testing.T, arabicNumber int, want string) {
	t.Helper()
	got := ConvertToRomanNumber(arabicNumber)

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
