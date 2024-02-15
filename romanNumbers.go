package main

import "errors"

func ConvertToRomanNumber(arabicNumber int) (string, error) {
	if arabicNumber <= 0 || arabicNumber > 3999 {
		return "", errors.New("arabic number must be between 1 and 3999")
	}

	arabicToRoman := []struct {
		arabic int
		roman  string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	result := ""

	for _, current := range arabicToRoman {
		for arabicNumber >= current.arabic {
			result += current.roman
			arabicNumber -= current.arabic
		}
	}

	return result, nil
}
