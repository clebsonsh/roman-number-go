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
