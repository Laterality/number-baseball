package kr.latera.baseball.domain;

import java.util.Arrays;
import java.util.Collection;
import java.util.HashSet;
import java.util.Set;
import java.util.stream.Collectors;

import static kr.latera.baseball.domain.BaseballNumbers.NUMS_LENGTH;

public class BaseballEngine {

    private final NumberGenerator generator;
    private BaseballNumbers answer;

    public BaseballEngine(NumberGenerator numberGenerator) {
        this.generator = numberGenerator;
    }

    public void start() {
        Set<Integer> nums = new HashSet<>();
        while (true) {
            nums.add(generator.generate());
            if (nums.size() == NUMS_LENGTH) {
                break;
            }
        }
        this.answer = new BaseballNumbers(nums);
    }

    public CheckResult check(int... input) {
        return check(Arrays.stream(input).boxed().collect(Collectors.toList()));
    }

    public CheckResult check(Collection<Integer> input) {
        BaseballNumbers other = new BaseballNumbers(input);
        return CheckResult.from(answer.countStrike(other), answer.countBall(other));
    }
}
