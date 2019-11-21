package kr.latera.baseball.domain;

import java.util.Arrays;
import java.util.LinkedList;
import java.util.Queue;

public class TestNumberGenerator implements NumberGenerator {

    private Queue<Integer> numsToGenerate;

    public TestNumberGenerator(int... numbers) {
        numsToGenerate = new LinkedList<>();
        Arrays.stream(numbers)
            .forEach(numsToGenerate::add);
    }

    @Override
    public int generate() {
        return numsToGenerate.poll();
    }
}
