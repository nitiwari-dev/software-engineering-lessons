package com.coderconsole.dsa.easy.hashmap

import org.assertj.core.api.Assertions
import org.junit.jupiter.api.BeforeEach
import org.junit.jupiter.api.DisplayName
import org.junit.jupiter.api.Test
import java.util.List

internal class AnagramsTest {
    private var anagrams: Anagrams? = null
    @BeforeEach
    fun setUp() {
        anagrams = Anagrams()
    }

    @Test
    @DisplayName("Given empty list of anagrams return empty list")
    fun groupAnagram_emptyAnagramList_returnEmpty() {
        val input = ArrayList<String>()
        val output = anagrams!!.groupAnagramBySort(input)
        Assertions.assertThat(output).isEmpty()
    }

    @Test
    @DisplayName("Given list of anagrams words group them together using sort")
    fun groupAnagram_listOfAnagrams_groupAnagramsUsingSorting() {
        val input = List.of("eat", "tea", "tan", "ate", "nat", "bat")
        val groups = List.of(List.of("eat", "tea", "ate"), List.of("tan", "nat"), List.of("bat"))
        val output = anagrams!!.groupAnagramBySort(input)
        Assertions.assertThat(output).hasSize(groups.size).hasSameElementsAs(groups)
    }

    @Test
    @DisplayName("Given list of anagrams words group them using character count")
    fun groupAnagram_listOfAnagrams_groupAnagramsByCharacterCount() {
        val input = List.of("eat", "tea", "tan", "ate", "nat", "bat")
        val groups = List.of(List.of("eat", "tea", "ate"), List.of("tan", "nat"), List.of("bat"))
        val output = anagrams!!.groupAnagramByCharacterCount(input)
        Assertions.assertThat(output).hasSize(groups.size).hasSameElementsAs(groups)
    }
}