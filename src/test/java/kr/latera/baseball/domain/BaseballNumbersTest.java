package kr.latera.baseball.domain;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertDoesNotThrow;
import static org.junit.jupiter.api.Assertions.assertThrows;

public class BaseballNumbersTest {

    @Test
    void create() {
        assertDoesNotThrow(() -> new BaseballNumbers(1, 2, 3));
    }

    @Test
    void duplicate() {
        assertThrows(InvalidNumbersException.class, () ->
                new BaseballNumbers(1, 1, 2));
    }
}
