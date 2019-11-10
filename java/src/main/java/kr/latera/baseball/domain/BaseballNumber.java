package kr.latera.baseball.domain;

import java.util.Objects;

public class BaseballNumber {

    static final int MAX_EXCLUSIVE = 10;
    static final int MIN_INCLUSIVE = 1;

    private final int value;

    private BaseballNumber(int value) {
        this.value = value;

        checkRange(value);
    }

    private void checkRange(int value) {
        if (value < MIN_INCLUSIVE || value >= MAX_EXCLUSIVE) {
            throw new InvalidNumberException(String.format("Number must be greater than or equal to %d and less than %d", MIN_INCLUSIVE, MAX_EXCLUSIVE));
        }
    }

    static BaseballNumber from(int value) {
        return new BaseballNumber(value);
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        BaseballNumber that = (BaseballNumber) o;
        return value == that.value;
    }

    @Override
    public int hashCode() {
        return Objects.hash(value);
    }

    @Override
    public String toString() {
        return "BaseballNumber{" +
                "value=" + value +
                '}';
    }
}
