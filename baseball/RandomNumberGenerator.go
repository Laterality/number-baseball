package baseball

import (
	"math/rand"
)

// RandomNumberGenerator implements NumberGenerator using ranom module
type randomNumberGenerator struct {
	min int
	max int
}

// Generate returns random number [min, max]
func (G *randomNumberGenerator) Generate() int {
	return rand.Intn(G.max-G.min) + G.min
}

// NewRandomNumberGenerator is factory method for number random number generator
func NewRandomNumberGenerator() NumberGenerator {
	generator := new(randomNumberGenerator)
	generator.min = NumberMin
	generator.max = NumberMax
	return generator
}
