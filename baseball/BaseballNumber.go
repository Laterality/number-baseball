package baseball

import (
	"fmt"
)

// Number is abstraction of single baseball number
type number = int

// Numbers is abstraction of group of baeball numbers
type Numbers interface {
	countStrike(other Numbers) int
	countBall(other Numbers) int
	visit(func(idx int, num number))
}

type numbers struct {
	nums []number
}

const (
	// NumberLength is length of input number slices
	NumberLength = 3
	// NumberMin is the minimum value of baseball number
	NumberMin = 1
	// NumberMax is the maximum value of baseball number
	NumberMax = 9
)

func (N numbers) countStrike(other Numbers) (strike int) {
	strike = 0
	other.visit(func(idx int, num number) {
		if N.nums[idx] == num {
			strike++
		}
	})
	return
}

func (N numbers) countBall(other Numbers) (ball int) {
	ball = 0
	for idx, num := range N.nums {
		other.visit(func(otherIdx int, otherNum number) {
			if idx != otherIdx && num == otherNum {
				ball++
			}
		})
	}
	return
}

func (N numbers) visit(visitor func(int, number)) {
	for idx, num := range N.nums {
		visitor(idx, num)
	}
}

// InvalidBaseballNumberError is for when trying to create baseball number with invalid value
type InvalidBaseballNumberError struct {
	cause int
}

func (E InvalidBaseballNumberError) Error() string {
	return fmt.Sprintf("Invalid number: %d", E.cause)
}

// InvalidBaseballNumbersLengthError is for when length of numbers are incorrect with rule of the game
type InvalidBaseballNumbersLengthError struct {
	actualLength int
}

func (E InvalidBaseballNumbersLengthError) Error() string {
	return fmt.Sprintf("Invalid length of numbers: %d", E.actualLength)
}

// DuplicateBaseballNumberError is for when list of numbers has numbers which duplicate
type DuplicateBaseballNumberError struct{}

func (E DuplicateBaseballNumberError) Error() string {
	return "숫자는 중복될 수 없습니다."
}

func newBaseballNumber(num int) (number, error) {
	if num < NumberMin || num > NumberMax {
		return 0, InvalidBaseballNumberError{num}
	}
	var created number = num
	return created, nil
}

// NewBaseballNumbers creates new group of baseball numbers
func newBaseballNumbers(nums []int) (Numbers, error) {
	if len(nums) != NumberLength {
		return nil, InvalidBaseballNumbersLengthError{len(nums)}
	}
	baseballNums := make([]number, 0, len(nums))
	set := make(map[int]int)
	for _, num := range nums {
		set[num]++
		created, err := newBaseballNumber(num)
		if err != nil {
			return nil, err
		}
		baseballNums = append(baseballNums, created)
	}
	if len(set) != NumberLength {
		return nil, DuplicateBaseballNumberError{}
	}
	return numbers{baseballNums}, nil
}
