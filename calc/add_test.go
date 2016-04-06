package calc

import "testing"

// 引き算のテストその１
func TestAdd(t *testing.T) {
	const x, y, expected = 1, 2, 3
	result := Add(x, y)
	if result != expected {
		t.Errorf("Add(%v, %v) = %v, want: %v", x, y, result, expected)
	}
}

// 引き算のテストその１
func TestAdd2(t *testing.T) {
	const x, y, expected = 2, 4, 6
	result := Add(x, y)
	if result != expected {
		t.Errorf("Add(%v, %v) = %v, want: %v", x, y, result, expected)
	}
}
