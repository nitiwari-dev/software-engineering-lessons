package com.coderconsole.dsa.easy.hashmap;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

import java.util.List;

import static org.assertj.core.api.AssertionsForClassTypes.assertThat;

class TargetPairTest {

    private TargetPair targetPair;

    @BeforeEach
    void setUp(){
        targetPair = new TargetPair();
    }

    @Test
    @DisplayName("Given target and list of numbers return true if the pair exists")
    void targetPair_pairExists_returnTrue(){
        var target = 10;
        var input = List.of(3, 4 ,5, 5);
        var output = targetPair.isPairExistsByHashMap(input, target);
        assertThat(output).isTrue();
    }

    @Test
    @DisplayName("Given target and list of numbers return false if the pair donot exists")
    void targetPair_pairDoNotExists_returnFalse(){
        var target = 10;
        var input = List.of(3, 4 ,5, 1);
        var output = targetPair.isPairExistsByHashMap(input, target);
        assertThat(output).isFalse();
    }

    @Test
    @DisplayName("Given target and list of numbers return false if the pair donot exists using brute force")
    void targetPair_pairDoNotExists_bruteForce_returnFalse(){
        var target = 12;
        var input = List.of(3, 4 ,5, 1);
        var output = targetPair.isPairExistsByBruteForce(input, target);
        assertThat(output).isFalse();
    }

    @Test
    @DisplayName("Given target and list of numbers return true if the pair exists using brute force")
    void targetPair_pairDoNotExists_bruteForce_returnTrue(){
        var target = 4;
        var input = List.of(3, 4 ,5, 1);
        var output = targetPair.isPairExistsByBruteForce(input, target);
        assertThat(output).isTrue();
    }
}
