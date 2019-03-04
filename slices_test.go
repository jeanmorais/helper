package helper

import (
	"reflect"
	"testing"
)

func TestContainsWithContainedString(t *testing.T) {

	result := Contains([]string{"A", "B", "C", "D", "E"}, "C")
	if !result {
		t.Errorf("Test failed. Expected %v, but returned %v", true, result)
	}
}

func TestContainsWithNotContainedString(t *testing.T) {

	result := Contains([]string{"A", "B", "C", "D", "E"}, "X")
	if result {
		t.Errorf("Test failed. Expected %v, but returned %v", true, result)
	}
}

func TestRemoveOrdered(t *testing.T) {

	tables := []struct {
		input    []string
		index    int
		expected []string
	}{
		{[]string{"A", "B", "C", "D", "E"}, 3, []string{"A", "B", "C", "E"}},
		{[]string{"A", "B", "C", "D", "E"}, 0, []string{"B", "C", "D", "E"}},
	}

	for _, table := range tables {
		result := RemoveOrdered(table.index, table.input)
		if !reflect.DeepEqual(result, table.expected) {
			t.Errorf("Test failed. Expected %v, but returned %v", table.expected, result)
		}
	}
}
