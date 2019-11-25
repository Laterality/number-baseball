package baseball

// Engine creates answer and proceeds client inputs
type Engine interface {
	// start initializes game state
	Start()
	// input checks number of strike and ball with argument
	Input([]int) (int, int, error)
}

// NumberGenerator is interface for generate number to make answer
type NumberGenerator interface {
	Generate() int
}

type baseballEngine struct {
	generator NumberGenerator
	answer    numbers
}

// CreateBaseballEngine constructs instance of baeball engine implementation
func CreateBaseballEngine(generator NumberGenerator) Engine {
	engine := new(baseballEngine)
	engine.generator = generator
	return engine
}

func (E *baseballEngine) Start() {
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

func (E *baseballEngine) Input(nums []int) (strike int, ball int, err error) {
	input, err := newBaseballNumbers(nums)
	if err != nil {
		return
	}
	strike = E.answer.countStrike(input)
	ball = E.answer.countBall(input)

	return
}
