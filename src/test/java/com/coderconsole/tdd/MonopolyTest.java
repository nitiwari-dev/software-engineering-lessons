package com.coderconsole.tdd;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;
import org.mockito.Mockito;

import static org.assertj.core.api.Assertions.assertThat;
import static org.mockito.BDDMockito.given;
import static org.mockito.Mockito.times;
import static org.mockito.Mockito.verify;

class MonopolyTest {

    private Monopoly spy;


    @BeforeEach
    public void setUp() {
        spy = Mockito.spy(new Monopoly(0)); //
    }

    //Naming Convention:

    //Given when then
    //when then
    //given when
    //returnWhen
    //doReturnWhen

    @Test
    @DisplayName("if the dice has different number return there new position")
    void givenDiceNumberReturnNewPosition() {
        given(spy.throwDice()).willReturn(1, 3);
        assertThat(spy.advance()).isEqualTo(4);
        verify(spy, times(2)).throwDice();
    }

    @Test
    @DisplayName("if the dice is 1 return current position")
    void givenDiceIsOneReturnCurrentPosition() {
        given(spy.throwDice()).willReturn(1);
        assertThat(spy.advance()).isEqualTo(0);
        verify(spy, times(6)).throwDice();
    }

    @Test
    @DisplayName("if the dice values are equal and reaches maximum repetition reset position to zero")
    void advanceDiceAndReturnPosition() {
        given(spy.throwDice()).willReturn(1, 1, 1, 1, 1, 1); //
        assertThat(spy.advance(3)).isZero();
        verify(spy, Mockito.atMost(6)).throwDice();
    }

    @Test
    @DisplayName("if the dice values are equal advance the position and stop when max limit is reached")
    void advanceDiceAndStopRecursion() {
        given(spy.throwDice()).willReturn(1, 1, 2, 3); //
        assertThat(spy.advance(3)).isEqualTo(7);
        verify(spy, times(4)).throwDice();
    }
}