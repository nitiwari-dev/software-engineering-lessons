package com.coderconsole.dsa.easy.hashmap

/**
 * Given an array of strings strs, group the anagrams together. You can return the answer in any order.
 *
 *
 * An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.
 *
 *
 * Example 1:
 *
 *
 * Input: strs = ["eat","tea","tan","ate","nat","bat"]
 * Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
 * Example 2:
 *
 *
 * Input: strs = [""]
 * Output: [[""]]
 * Example 3:
 *
 *
 * Input: strs = ["a"]
 * Output: [["a"]]
 *
 *
 *
 *
 * Approach 1 using: Sorting
 * T: O(N*MLogM)
 * S: O(N*M) -> N is number of elements in the array and M is the maximum length of string in array
 *
 *
 * Approach 2 using: Character count
 * T: O(N*M)
 * S: O(N*M) -> N is number of elements in the array and M is the maximum length of string in array
 */
class Anagrams {
    fun groupAnagramBySort(input: List<String>): List<List<String>> {
        val map = HashMap<String, MutableList<String>>()
        for (originalKey in input) {
            val sortedKey = getSortedKey(originalKey)
            val group = map.getOrDefault(sortedKey, ArrayList())
            group.add(originalKey)
            map[sortedKey] = group
        }
        return ArrayList<List<String>>(map.values)
    }

    fun groupAnagramByCharacterCount(input: List<String>): List<List<String>> {
        val map = HashMap<String, MutableList<String>>()
        val charInput = IntArray(26)
        for (originalKey in input) {
            charInput.fill(0)
            val key = getCharacterKey(originalKey, charInput)
            val group = map.getOrDefault(key, ArrayList())
            group.add(originalKey)
            map[key] = group
        }
        return ArrayList<List<String>>(map.values)
    }

    private fun getSortedKey(originalKey: String): String {
        val chArray = originalKey.toCharArray().apply { sort() }
        return String(chArray)
    }

    private fun getCharacterKey(originalKey: String, characterCount: IntArray): String {
        for (ch in originalKey.toCharArray()) {
            characterCount[ch.code - 'a'.code] += 1
        }
        val builder = StringBuilder()
        for (count in characterCount) {
            builder.append("#")
            builder.append(count)
        }
        return builder.toString()
    }
}