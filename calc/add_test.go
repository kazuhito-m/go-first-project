package calc

import "testing"

func Test足し算のテストその01(t *testing.T) {
	const x, y, expected = 1, 2, 3
	actual := Add(x, y)
	if actual != expected {
		t.Errorf("Add(%v, %v) = %v, want: %v", x, y, actual, expected)
	}
}

func Test足し算のテストその02(t *testing.T) {
	const x, y, expected = 2, 4, 6
	actual := Add(x, y)
	if actual != expected {
		t.Errorf("Add(%v, %v) = %v, want: %v", x, y, actual, expected)
	}
}
