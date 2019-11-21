package kr.latera.baseball.domain;

public class InvalidNumberException extends RuntimeException{

    public InvalidNumberException(String msg) {
        super(msg);
    }
}
