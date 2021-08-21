package com.coderconsole.tdd;

import java.util.Random;

public class Monopoly {
    private static final int MAX = 3;
    private int currentPosition; // startPosition

    public Monopoly(int startingPosition) {
        currentPosition = startingPosition;
    }

    public int throwDice() {
        Random rand = new Random();
        return rand.nextInt(5) + 1;
    }

    public int advance(){
      return advance(MAX);
    }


    public int advance(int count) { // base condition
        if (count == 0) {
            currentPosition = 0;
            return currentPosition;
        }

        int dice1 = throwDice(); // 2 0
        int dice2 = throwDice(); // 3 1

        currentPosition += dice1 + dice2;
        if (dice1 == dice2) {
            return advance(count - 1);
        }
        return currentPosition;


    }
}
