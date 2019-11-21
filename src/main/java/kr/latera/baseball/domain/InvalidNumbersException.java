package kr.latera.baseball.domain;

public class InvalidNumbersException extends RuntimeException {
    public InvalidNumbersException(String message) {
        super(message);
    }

    public InvalidNumbersException(Throwable cause) {
        super(cause);
    }
}
