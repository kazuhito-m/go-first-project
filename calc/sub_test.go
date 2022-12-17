package calc

import "testing"

func Test引き算のテストその01(t *testing.T) {
	const x, y, expected = 3, 2, 1
	actual := Sub(x, y)
	if actual != expected {
		t.Errorf("sub(%v, %v) = %v, want: %v", x, y, actual, expected)
	}
}

func Test引き算のテストその02(t *testing.T) {
	const x, y, expected = 6, 4, 2
	actual := Sub(x, y)
	if actual != expected {
		t.Errorf("sub(%v, %v) = %v, want: %v", x, y, actual, expected)
	}
}
