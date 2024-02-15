package main

func ConvertToRomanNumber(arabicNumber int) string {
	arabicToRoman := []struct {
		arabic int
		roman  string
	}{
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

	return result
}
