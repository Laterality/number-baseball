package baseball

// Engine creates answer and proceeds client inputs
type Engine interface {
	// start initializes game state
	start()
	// input checks number of strike and ball with argument
	input([]int) (int, int)
}

// NumberGenerator is interface for generate number to make answer
type NumberGenerator interface {
	Generate() int
}

type baseballEngine struct {
	generator NumberGenerator
	answer    numbers
}

type numbers struct {
	nums []int
}

const (
	// NumberLength is length of input number slices
	NumberLength = 3
	// NumberMin is the minimum value of baseball number
	NumberMin = 1
	// NumberMax is the maximum value of baseball number
	NumberMax = 9
)

// CreateBaseballEngine constructs instance of baeball engine implementation
func CreateBaseballEngine(generator NumberGenerator) Engine {
	engine := new(baseballEngine)
	engine.generator = generator
	return engine
}

func (E *baseballEngine) start() {
	E.answer.nums = make([]int, NumberLength)
	nums := &E.answer.nums
	idx := 0
	for idx < len(*nums) {
		num := E.generator.Generate()
		if indexOf(*nums, num) == -1 {
			(*nums)[idx] = num
			idx++
		}
	}
}

func indexOf(nums []int, target int) int {
	for idx, num := range nums {
		if num == target {
			return idx
		}
	}
	return -1
}

func (E *baseballEngine) input(nums []int) (strike int, ball int) {
	input := numbers{nums}
	strike = E.answer.countStrike(&input)
	ball = E.answer.countBall(&input)

	return
}

func (N *numbers) countStrike(other *numbers) (strike int) {
	strike = 0
	for idx := range other.nums {
		if N.nums[idx] == other.nums[idx] {
			strike++
		}
	}
	return
}

func (N *numbers) countBall(other *numbers) (ball int) {
	ball = 0
	for idx, num := range N.nums {
		for otherIdx, otherNum := range other.nums {
			if idx != otherIdx && num == otherNum {
				ball++
			}
		}
	}
	return
}
