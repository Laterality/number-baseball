package kr.latera.baseball.domain;

import org.junit.jupiter.api.Test;

import java.util.List;
import java.util.stream.Collectors;
import java.util.stream.IntStream;

import static org.assertj.core.api.Assertions.assertThat;
import static org.junit.jupiter.api.Assertions.assertDoesNotThrow;

public class RandomNumberGeneratorTest {

    @Test
    void generate() {
        assertDoesNotThrow(() -> new RandomNumberGenerator(1, 100));
    }

    @Test
    void generate_boundary() {
        final int min = 1;
        final int max = 10;
        NumberGenerator generator = new RandomNumberGenerator(min, max);
        List<Integer> generated = IntStream.range(1, 1000)
                .mapToObj(i -> generator.generate())
                .collect(Collectors.toList());
        assertThat(generated)
                .contains(min)
                .contains(max - 1)
                .doesNotContain(max)
                .doesNotContain(min - 1);
    }
}
