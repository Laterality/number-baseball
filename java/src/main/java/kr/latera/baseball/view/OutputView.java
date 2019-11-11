package kr.latera.baseball.view;

import java.io.PrintStream;

public class OutputView {

    private final PrintStream printStream;

    public OutputView(PrintStream printStream) {
        this.printStream = printStream;
    }

    public void promptInputNumbers(String separator) {
        printStream.println(String.format("숫자를 입력하세요(%s로 구분)", separator));
    }

    public void printCheckResult(int strike, int ball) {
        printStream.println(String.format("%d 스트라이크, %d 볼", strike, ball));
    }

    public void printEndWithWin() {
        printStream.println("정답입니다!");
    }

    public void printError(Throwable e) {
        printStream.println(e.getMessage());
    }
}
