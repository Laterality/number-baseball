package kr.latera.baseball.domain;

public class BaseballEngineFactory {

    public static BaseballEngine createBaseballEngine() {
        return new BaseballEngine(new RandomNumberGenerator(BaseballNumber.MIN_INCLUSIVE, BaseballNumber.MAX_EXCLUSIVE));
    }
}
