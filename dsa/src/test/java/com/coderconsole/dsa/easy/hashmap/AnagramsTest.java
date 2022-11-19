package com.coderconsole.dsa.easy.hashmap;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

import java.util.List;

import static org.assertj.core.api.Assertions.assertThat;

class AnagramsTest {

    Anagrams anagrams;

    @BeforeEach
    void setUp() {
        anagrams = new Anagrams();
    }


    @Test
    @DisplayName("Given list of anagrams words group them together and return list of each group")
    void returnGroupOfAnagramsUsingSorting() {
        var input = List.of("eat","tea","tan","ate","nat","bat");
        var output = anagrams.groupAnagramTogetherWithSorting(input);
        assertThat(output).isNotEmpty();
    }

}