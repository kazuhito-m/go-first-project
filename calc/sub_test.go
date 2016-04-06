package calc

import "testing"

// 引き算のテストその１
func TestSub(t *testing.T) {
	const x, y, expected = 3, 2, 1
	actual := Sub(x, y)
	if actual != expected {
		t.Errorf("sub(%v, %v) = %v, want: %v", x, y, actual, expected)
	}
}

// 引き算のテストその１
func TestSub2(t *testing.T) {
	const x, y, expected = 6, 4, 2
	actual := Sub(x, y)
	if actual != expected {
		t.Errorf("sub(%v, %v) = %v, want: %v", x, y, actual, expected)
	}
}
