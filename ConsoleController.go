package main

import "github.com/laterality/number-baseball/baseball"

// GameSession is abstraction of single game session
type GameSession interface {
	proceed()
	isFinished() bool
}

type gameSession struct {
	engine   baseball.Engine
	input    ConsoleInputView
	output   ConsoleOutputView
	finished bool
}

func (S *gameSession) proceed() {
	nums := S.input.promptBaseballNumbers(baseball.NumberMin, baseball.NumberMax, baseball.NumberLength)
	strike, ball, err := S.engine.Input(nums)
	if err != nil {
		S.output.printError(err.Error())
		return
	}

	if strike == baseball.NumberLength {
		S.output.printSuccess()
		S.finished = true
		return
	}

	S.output.printResult(strike, ball)
}

func (S *gameSession) isFinished() bool {
	return S.finished
}

// NewGameSession constructs new game session
func NewGameSession() GameSession {
	sess := new(gameSession)
	sess.engine = baseball.CreateBaseballEngine(baseball.NewRandomNumberGenerator())
	sess.input = ConsoleInputView{}
	sess.output = ConsoleOutputView{}
	sess.finished = false
	sess.engine.Start()
	return sess
}
