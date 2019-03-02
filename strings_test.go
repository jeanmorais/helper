package gpad

import "testing"

func TestOnlyLettersWithNumbers(t *testing.T) {

	expected := ""
	formatted := OnlyLetters("123")
	if formatted != expected {
		t.Errorf("Test failed. Expected %v, but returned %v", expected, formatted)
	}
}

func TestOnlyLettersWithNumbersAndLettersAndSpecialChars(t *testing.T) {

	expected := "golang"
	formatted := OnlyLetters("1gol@ang2.3!")
	if formatted != expected {
		t.Errorf("Test failed. Expected %v, but returned %v", expected, formatted)
	}
}

func TestOnlyNumbersWithLetters(t *testing.T) {

	expected := ""
	formatted := OnlyNumbers("golang")
	if formatted != expected {
		t.Errorf("Test failed. Expected %v, but returned %v", expected, formatted)
	}
}

func TestOnlyNumbersWithLettersAndNumbersAndSpecialChars(t *testing.T) {

	expected := "123"
	formatted := OnlyNumbers("1gol@ang2.3!")
	if formatted != expected {
		t.Errorf("Test failed. Expected %v, but returned %v", expected, formatted)
	}
}