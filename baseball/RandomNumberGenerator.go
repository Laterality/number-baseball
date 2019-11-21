package baseball

import "math/rand"

// RandomNumberGenerator implements NumberGenerator using ranom module
type RandomNumberGenerator struct {
	min int
	max int
}

// Generate returns random number [min, max]
func (G *RandomNumberGenerator) Generate() int {
	return rand.Intn(G.max-G.min) + G.min
}
