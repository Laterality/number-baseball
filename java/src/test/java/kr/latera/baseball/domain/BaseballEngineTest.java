package kr.latera.baseball.domain;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;

import static org.assertj.core.api.Assertions.assertThat;
import static org.junit.jupiter.api.Assertions.assertDoesNotThrow;
import static org.junit.jupiter.api.Assertions.assertThrows;

public class BaseballEngineTest {

    private BaseballEngine engine;

    @BeforeEach
    void init() {
        engine = new BaseballEngine(new TestNumberGenerator(1, 2, 3));
    }

    @Test
    void create() {
        assertDoesNotThrow(() -> new BaseballEngine(new TestNumberGenerator(1, 2, 3)));
    }

    @Test
    void create_random() {
        for (int i = 0; i < 100; i++) {
            assertDoesNotThrow(() -> new BaseballEngine(
                    new RandomNumberGenerator(
                            BaseballNumber.MIN_INCLUSIVE,
                            BaseballNumber.MAX_EXCLUSIVE))).start();
        }
    }

    @Test
    void succeed() {
        engine.start();
        CheckResult result = engine.check(1, 2, 3);
        assertThat(result.isCorrect()).isTrue();
    }

    @Test
    void ball_3() {
        engine.start();
        CheckResult result = engine.check(2, 3, 1);
        assertThat(result.getBall()).isEqualTo(3);
    }

    @Test
    void result_mixed() {
        engine.start();
        CheckResult result = engine.check(2, 1, 3);
        assertThat(result.getStrike()).isEqualTo(1);
        assertThat(result.getBall()).isEqualTo(2);
    }

    @Test
    void number_boundary_min() {
        assertThrows(InvalidNumberException.class, () -> {
            engine.start();
            engine.check(0, 1, 2);
        });
    }
}
