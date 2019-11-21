package baseball

import "testing"

// TestNumberGenerator implements NumberGenerator only for test
type TestNumberGenerator struct {
	nums    []int
	current int
}

// Generate returns nums
func (T *TestNumberGenerator) Generate() (num int) {
	num = T.nums[T.current]
	T.current++
	return
}

// NewTestNumberGenerator constructs test number generator
func NewTestNumberGenerator(nums []int) *TestNumberGenerator {
	generator := new(TestNumberGenerator)
	generator.current = 0
	generator.nums = nums
	return generator
}

func TestGame(t *testing.T) {
	engine := CreateBaseballEngine(NewTestNumberGenerator([]int{1, 2, 3}))
	engine.start()
	strike, ball := engine.input([]int{1, 3, 2})
	strikeExpected := 1
	ballExpected := 2
	if strike != strikeExpected {
		t.Errorf("Strike is incorrect, expected %d, but actual was %d", strikeExpected, strike)
	}
	if ball != ballExpected {
		t.Errorf("Ball is not incorrect, expected %d, but actual was %d", ballExpected, ball)
	}
}

func TestNoDuplicateNumberWithinAnswer(t *testing.T) {
	engine := CreateBaseballEngine(NewTestNumberGenerator([]int{1, 1, 2, 2, 2, 1, 4}))
	engine.start()
	strike, _ := engine.input([]int{1, 2, 4})
	if strike != 3 {
		t.Error("Duplicate number is in answer")
	}
}
