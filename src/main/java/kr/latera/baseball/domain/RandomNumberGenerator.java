package kr.latera.baseball.domain;

import java.util.Random;

public class RandomNumberGenerator implements NumberGenerator {

    private final Random random;
    private final int min;
    private final int max;

    public RandomNumberGenerator(int minInclusive, int maxExclusive) {
        random = new Random();
        this.min = minInclusive;
        this.max = maxExclusive;
    }

    @Override
    public int generate() {
        return random.nextInt(max - min) + min;
    }
}
