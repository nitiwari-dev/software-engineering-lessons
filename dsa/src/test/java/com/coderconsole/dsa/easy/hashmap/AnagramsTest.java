package com.coderconsole.dsa.easy.hashmap;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

import java.util.ArrayList;
import java.util.List;

import static org.assertj.core.api.Assertions.assertThat;

class AnagramsTest {

    private Anagrams anagrams;

    @BeforeEach
    void setUp() {
        anagrams = new Anagrams();
    }


    @Test
    @DisplayName("Given empty list of anagrams return empty list")
    void groupAnagram_emptyAnagramList_returnEmpty() {
        var input = new ArrayList<String>();
        var output = anagrams.groupAnagramWithSort(input);
        assertThat(output).isEmpty();
    }

    @Test
    @DisplayName("Given list of anagrams words group them together and return list of each group")
    void groupAnagram_listOfAnagrams_groupAnagramsUsingSorting() {
        var input = List.of("eat", "tea", "tan", "ate", "nat", "bat");
        var groups = List.of(List.of("eat", "tea", "ate"), List.of("tan", "nat"), List.of("bat"));
        var output = anagrams.groupAnagramWithSort(input);
        assertThat(output).hasSize(groups.size()).hasSameElementsAs(groups);
    }
}