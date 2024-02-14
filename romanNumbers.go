package main

func ConvertToRomanNumber(arabicNumber int) string {
	arabicToRoman := map[int]string{
		5: "V",
		4: "IV",
		1: "I",
	}

	result := ""

	for arabic, roman := range arabicToRoman {
		for arabicNumber >= arabic {
			result += roman
			arabicNumber -= arabic
		}
	}

	return result
}
