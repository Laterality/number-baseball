package baseball

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	min := 2
	max := 10
	loop := 1000
	generator := RandomNumberGenerator{min: min, max: max}
	minPresent := false
	maxPresent := false

	for i := 0; i < loop; i++ {
		num := generator.Generate()
		if num < min || num >= max {
			t.Errorf("Invalid number generated %d", num)
		}
		if num == min {
			minPresent = true
		}
		if num == max-1 {
			maxPresent = true
		}
	}

	if !minPresent {
		t.Error("Minimum boundary value is not generated")
	}

	if !maxPresent {
		t.Error("Maxium boundary value is not generated")
	}
}
