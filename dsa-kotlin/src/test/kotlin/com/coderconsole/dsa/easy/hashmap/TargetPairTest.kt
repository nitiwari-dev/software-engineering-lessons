package com.coderconsole.dsa.easy.hashmap

import org.assertj.core.api.Assertions.assertThat
import org.junit.jupiter.api.BeforeAll
import org.junit.jupiter.api.DisplayName
import org.junit.jupiter.api.Test
import org.junit.jupiter.api.TestInstance


@TestInstance(TestInstance.Lifecycle.PER_CLASS)
internal class TargetPairTest {

    private lateinit var targetPair: TargetPair

    @BeforeAll
    fun setUp() {
        targetPair = TargetPair()
    }

    @Test
    @DisplayName("Given target and list of numbers return true if the pair exists")
    fun targetPair_pairExists_returnTrue() {
        val target = 10
        val input = listOf(3, 4, 5, 5)
        val output = targetPair.isPairExistsByHashMap(input, target)
        assertThat(output).isTrue
    }

    @Test
    @DisplayName("Given target and list of numbers return false if the pair donot exists")
    fun targetPair_pairDoNotExists_returnFalse() {
        val target = 10
        val input = listOf(3, 4, 5, 1)
        val output = targetPair.isPairExistsByHashMap(input, target)
        assertThat(output).isFalse
    }

    @Test
    @DisplayName("Given target and list of numbers return false if the pair donot exists using brute force")
    fun targetPair_pairDoNotExists_bruteForce_returnFalse() {
        val target = 12
        val input = listOf(3, 4, 5, 1)
        val output = targetPair.isPairExistsByBruteForce(input, target)
        assertThat(output).isFalse
    }

    @Test
    @DisplayName("Given target and list of numbers return true if the pair exists using brute force")
    fun targetPair_pairDoNotExists_bruteForce_returnTrue() {
        val target = 4
        val input = listOf(3, 4, 5, 1)
        val output = targetPair.isPairExistsByBruteForce(input, target)
        assertThat(output).isTrue
    }
}