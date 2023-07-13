package com.coderconsole.dsa.easy.hashmap

import org.assertj.core.api.Assertions.assertThat
import org.junit.jupiter.api.BeforeAll
import org.junit.jupiter.api.DisplayName
import org.junit.jupiter.api.Test
import org.junit.jupiter.api.TestInstance

@TestInstance(TestInstance.Lifecycle.PER_CLASS)
internal class AnagramsTest {
    private lateinit var anagrams: Anagrams
    @BeforeAll
    fun setUp() {
        anagrams = Anagrams()
    }

    @Test
    @DisplayName("Given empty list of anagrams return empty list")
    fun groupAnagram_emptyAnagramList_returnEmpty() {
        val input = ArrayList<String>()
        val output = anagrams.groupAnagramBySort(input)
        assertThat(output).isEmpty()
    }

    @Test
    @DisplayName("Given list of anagrams words group them together using sort")
    fun groupAnagram_listOfAnagrams_groupAnagramsUsingSorting() {
        val input = listOf("eat", "tea", "tan", "ate", "nat", "bat")
        val groups = listOf(listOf("eat", "tea", "ate"), listOf("tan", "nat"), listOf("bat"))
        val output = anagrams.groupAnagramBySort(input)
        assertThat(output).hasSize(groups.size).hasSameElementsAs(groups)
    }

    @Test
    @DisplayName("Given list of anagrams words group them using character count")
    fun groupAnagram_listOfAnagrams_groupAnagramsByCharacterCount() {
        val input = listOf("eat", "tea", "tan", "ate", "nat", "bat")
        val groups = listOf(listOf("eat", "tea", "ate"), listOf("tan", "nat"), listOf("bat"))
        val output = anagrams.groupAnagramByCharacterCount(input)
        assertThat(output).hasSize(groups.size).hasSameElementsAs(groups)
    }
}