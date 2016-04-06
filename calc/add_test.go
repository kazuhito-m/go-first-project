package calc

import "testing"

// 引き算のテストその１
func TestAdd(t *testing.T) {
	const x, y, expected = 1, 2, 3
	actual := Add(x, y)
	if actual != expected {
		t.Errorf("Add(%v, %v) = %v, want: %v", x, y, actual, expected)
	}
}

// 引き算のテストその１
func TestAdd2(t *testing.T) {
	const x, y, expected = 2, 4, 6
	actual := Add(x, y)
	if actual != expected {
		t.Errorf("Add(%v, %v) = %v, want: %v", x, y, actual, expected)
	}
}
