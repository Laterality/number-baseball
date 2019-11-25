package baseball

import (
	"testing"
)

func TestNewNumbers(t *testing.T) {
	_, err := newBaseballNumbers([]int{1, 2, 3})
	if err != nil {
		t.Errorf("Error is occurred while create baseball numbers")
	}
}
