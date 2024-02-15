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
		{1000, "M"},
		{2000, "MM"},
		{3000, "MMM"},
		{3999, "MMMCMXCIX"},
	}

	for _, tt := range numbersToTest {
		assertConvertToRomanNumber(t, tt.arabicNumber, tt.romanNumber)
	}

	t.Run("test with 0", func(t *testing.T) {
		_, err := ConvertToRomanNumber(0)

		assertError(t, err)
	})

	t.Run("test with negative number", func(t *testing.T) {
		_, err := ConvertToRomanNumber(0)

		assertError(t, err)
	})

	t.Run("test with number greater than 3999", func(t *testing.T) {
		_, err := ConvertToRomanNumber(0)

		assertError(t, err)
	})

}

func assertConvertToRomanNumber(t *testing.T, arabicNumber int, want string) {
	t.Helper()
	got, _ := ConvertToRomanNumber(arabicNumber)

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Errorf("expected an error but didn't get one")
	}
}
