package helper

import "testing"

var sliceTest []string = []string{"go", "lang"}

func TestContainsWithContainedString(t *testing.T) {

	result := Contains(sliceTest, "lang")
	if !result {
		t.Errorf("Test failed. Expected %v, but returned %v", true, result)
	}
}

func TestContainsWithNotContainedString(t *testing.T) {

	result := Contains(sliceTest, "long")
	if result {
		t.Errorf("Test failed. Expected %v, but returned %v", true, result)
	}
}
