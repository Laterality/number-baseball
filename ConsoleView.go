package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ConsoleInputView is view for prompt commandline input
type ConsoleInputView struct{}

func (V *ConsoleInputView) promptBaseballNumbers(min int, max int, length int) []int {
	fmt.Printf("%d부터 %d 까지의 숫자 %d 개를 쉼표로 구분하여 입력하세요.\n", min, max, length)
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	tokens := strings.Split(s, ",")
	trimSpace(&tokens)
	filterEmpty(&tokens)
	return toNumbers(&tokens)
}

func trimSpace(tokens *[]string) {
	for idx, str := range *tokens {
		(*tokens)[idx] = strings.TrimSpace(str)
	}
}

func filterEmpty(tokens *[]string) {
	for idx, token := range *tokens {
		if len(token) == 0 {
			*tokens = append((*tokens)[:idx], (*tokens)[idx+1:]...)
		}
	}
}

func toNumbers(tokens *[]string) (nums []int) {
	nums = make([]int, 0, len(*tokens))
	for _, str := range *tokens {
		n, _ := strconv.Atoi(str)
		nums = append(nums, n)
	}
	return
}

// ConsoleOutputView is view for print commandline output
type ConsoleOutputView struct{}

func (V *ConsoleOutputView) printResult(strike, ball int) {
	fmt.Printf("%d스트라이크, %d 볼\n", strike, ball)
}

func (V *ConsoleOutputView) printError(message string) {
	fmt.Println(message)
}

func (V *ConsoleOutputView) printSuccess() {
	fmt.Println("정답입니다!")
}
