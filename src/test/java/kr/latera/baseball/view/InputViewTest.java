package kr.latera.baseball.view;

import org.junit.jupiter.api.Test;

import java.io.ByteArrayInputStream;

import static org.assertj.core.api.Assertions.assertThat;

public class InputViewTest {

    @Test
    void input_numbers() {
        InputView iv = InputView.of(new ByteArrayInputStream("1, 2, 3".getBytes()));
        assertThat(iv.inputNumbers()).containsExactly(1, 2, 3);
    }
}
