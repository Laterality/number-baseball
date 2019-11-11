package kr.latera.baseball;

import kr.latera.baseball.domain.BaseballEngine;
import kr.latera.baseball.domain.BaseballEngineFactory;
import kr.latera.baseball.domain.CheckResult;
import kr.latera.baseball.view.InputView;
import kr.latera.baseball.view.OutputView;

public class ConsoleController {

    private final InputView inputView;
    private final OutputView outputView;
    private final BaseballEngine baseballEngine;

    public ConsoleController(InputView inputView, OutputView outputView) {
        this.inputView = inputView;
        this.outputView = outputView;
        baseballEngine = BaseballEngineFactory.createBaseballEngine();
    }

    public void run() {
        baseballEngine.start();

        while (true) {
            CheckResult result = promptAndCheckNumbers();
            if (result.isCorrect()) {
                outputView.printEndWithWin();
                break;
            }
            outputView.printCheckResult(result.getStrike(), result.getBall());
        }
    }

    private CheckResult promptAndCheckNumbers() {
        while (true) {
            try {
                outputView.promptInputNumbers(InputView.NUMBERS_SEPRATOR);
                return baseballEngine.check(inputView.inputNumbers());
            } catch (Exception e) {
                outputView.printError(e);
            }
        }
    }
}
