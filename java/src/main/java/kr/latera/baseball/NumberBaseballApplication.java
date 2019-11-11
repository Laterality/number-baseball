package kr.latera.baseball;

import kr.latera.baseball.view.InputView;
import kr.latera.baseball.view.OutputView;

public class NumberBaseballApplication {
    public static void main(String[] args) {
        new ConsoleController(InputView.of(System.in), new OutputView(System.out)).run();
    }
}
