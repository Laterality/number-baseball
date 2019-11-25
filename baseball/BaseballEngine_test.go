package baseball

import (
	"reflect"
	"testing"
)

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
	engine.Start()
	strike, ball, _ := engine.Input([]int{1, 3, 2})
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
	engine.Start()
	strike, _, _ := engine.Input([]int{1, 2, 4})
	if strike != 3 {
		t.Error("Duplicate number is in answer")
	}
}

func TestInvalidNumber(t *testing.T) {
	engine := CreateBaseballEngine(NewTestNumberGenerator([]int{1, 2, 3}))
	engine.Start()
	_, _, err := engine.Input([]int{0, 1, 2})
	if err == nil {
		t.Error("Error is not occurred on invalid number")
		return
	}
	if err.Error() != "Invalid number: 0" {
		t.Errorf("Another error occurred: %v", err)
	}
}

func TestInvalidLength(t *testing.T) {
	engine := CreateBaseballEngine(NewTestNumberGenerator([]int{1, 2, 3}))
	engine.Start()
	_, _, err := engine.Input([]int{1, 2})

	if err == nil {
		t.Error("Error is not occurred on invalid lnegth of numbers")
		return
	}

	if reflect.TypeOf(err) != reflect.TypeOf(*new(InvalidBaseballNumbersLengthError)) {
		t.Errorf("Another error occurred: %v", err)
	}
}

func TestDuplicateNumbers(t *testing.T) {
	engine := CreateBaseballEngine(NewTestNumberGenerator([]int{1, 2, 3}))
	engine.Start()
	_, _, err := engine.Input([]int{1, 2, 2})

	if err == nil {
		t.Error("Error is not occurred on invalid lnegth of numbers")
		return
	}

	if reflect.TypeOf(err) != reflect.TypeOf(*new(DuplicateBaseballNumberError)) {
		t.Errorf("Another error occurred: %v", err)
	}
}
