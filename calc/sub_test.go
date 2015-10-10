package calc

import "testing"


// 引き算のテストその１
func TestSub(t *testing.T) {
  const x, y, expected = 3, 2, 1
  result := Sub(x, y)
  if result != expected {
    t.Errorf("sub(%v, %v) = %v, want: %v", x, y, result, expected)
  }
}

// 引き算のテストその１
func TestSub2(t *testing.T) {
  const x, y, expected = 6, 4, 2
  result := Sub(x, y)
  if result != expected {
    t.Errorf("sub(%v, %v) = %v, want: %v", x, y, result, expected)
  }
}
