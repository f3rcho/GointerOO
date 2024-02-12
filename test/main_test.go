package main

import "testing"

func TestSum(t *testing.T) {
	result := Sum(2, 3)
	if result != 5 {
		t.Error("Expected 5, got ", result)
	}

	table := []struct {
		a        int
		b        int
		expected int
	}{
		{2, 3, 5},
		{5, 5, 10},
		{1, 2, 3},
	}
	for _, test := range table {
		result := Sum(test.a, test.b)
		if result != test.expected {
			t.Error("Expected:", test.expected, "got:", result)
		}
	}
}

func TestMax(t *testing.T) {
	table := []struct {
		a        int
		b        int
		expected int
	}{
		{2, 3, 3},
		{5, 6, 6},
		{1, 2, 2},
		{2, 1, 2},
	}

	for _, test := range table {
		result := GetMax(test.a, test.b)
		if result != test.expected {
			t.Error("Expected:", test.expected, "got:", result)
		}
	}
}
