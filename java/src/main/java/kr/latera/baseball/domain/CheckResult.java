package kr.latera.baseball.domain;

public class CheckResult {

    private final int strike;
    private final int ball;

    private CheckResult(int strike, int ball) {
        this.strike = strike;
        this.ball = ball;
    }

    static CheckResult from(int strike, int ball) {
        return new CheckResult(strike, ball);
    }

    public boolean isCorrect() {
        return strike == 3;
    }

    public int getStrike() {
        return strike;
    }

    public int getBall() {
        return ball;
    }
}
