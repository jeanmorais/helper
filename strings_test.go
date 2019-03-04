package helper

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

func TestRightPad(t *testing.T) {

	tables := []struct {
		str      string
		len      int
		pad      string
		expected string
	}{
		{"084970589", 11, "0", "08497058900"},
		{"golang", 8, ".", "golang.."},
		{"golang", 6, ".", "golang"},
	}

	for _, table := range tables {
		formatted := RightPad(table.str, table.len, table.pad)
		if formatted != table.expected {
			t.Errorf("Test failed. Expected %v, but returned %v", table.expected, formatted)
		}
	}
}

func TestLeftPad(t *testing.T) {

	tables := []struct {
		str      string
		len      int
		pad      string
		expected string
	}{
		{"084970589", 11, "0", "00084970589"},
		{"golang", 8, ".", "..golang"},
		{"golang", 6, ".", "golang"},
	}

	for _, table := range tables {
		formatted := LeftPad(table.str, table.len, table.pad)
		if formatted != table.expected {
			t.Errorf("Test failed. Expected %v, but returned %v", table.expected, formatted)
		}
	}
}

func TestReverse(t *testing.T) {

	tables := []struct {
		input    string
		expected string
	}{
		{"084970589", "985079480"},
		{"golang", "gnalog"},
	}

	for _, table := range tables {
		reverse := Reverse(table.input)
		if reverse != table.expected {
			t.Errorf("Test failed. Expected %v, but returned %v", table.expected, reverse)
		}
	}
}
