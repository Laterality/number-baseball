package kr.latera.baseball.domain;

import java.util.Arrays;
import java.util.Collection;
import java.util.HashSet;
import java.util.List;
import java.util.stream.Collectors;
import java.util.stream.IntStream;

public class BaseballNumbers {

    public static final int NUMS_LENGTH = 3;

    private final List<BaseballNumber> numbers;

    BaseballNumbers(int... numbers) {
       this(Arrays.stream(numbers)
               .boxed()
               .collect(Collectors.toList()));
    }

    BaseballNumbers(Collection<Integer> numbers) {
        this.numbers = numbers.stream()
                .map(BaseballNumber::from)
                .collect(Collectors.toList());
        ensureNumsLength();
    }

    private void ensureNumsLength() {
        int actualNumbersLength = new HashSet<>(this.numbers).size();
        if (actualNumbersLength != NUMS_LENGTH) {
            throw new InvalidNumbersException(String.format("숫자는 중복되지 않는 %d 개의 정수여야 합니다: %s", NUMS_LENGTH, numbers));
        }
    }

    int countStrike(BaseballNumbers other) {
        return (int) IntStream.range(0, NUMS_LENGTH)
                .filter(i -> numbers.get(i).equals(other.numbers.get(i)))
                .count();
    }

    int countBall(BaseballNumbers other) {
        List<BaseballNumber> otherNumbers = other.numbers;
        return (int) IntStream.range(0, NUMS_LENGTH)
                .filter(i -> numbers.indexOf(otherNumbers.get(i)) > -1)
                .filter(i -> !numbers.get(i).equals(otherNumbers.get(i)))
                .count();
    }
}
