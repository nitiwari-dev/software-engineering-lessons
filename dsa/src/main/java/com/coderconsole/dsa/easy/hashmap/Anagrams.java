package com.coderconsole.dsa.easy.hashmap;


import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashMap;
import java.util.List;

/**
 * Given an array of strings strs, group the anagrams together. You can return the answer in any order.
 * <p>
 * An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.
 * <p>
 * Example 1:
 * <p>
 * Input: strs = ["eat","tea","tan","ate","nat","bat"]
 * Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
 * Example 2:
 * <p>
 * Input: strs = [""]
 * Output: [[""]]
 * Example 3:
 * <p>
 * Input: strs = ["a"]
 * Output: [["a"]]
 * <p>
 * <p>
 * Approach 1 using: Sorting
 * T: O(N*MLogM)
 * S: O(N*M) -> N is number of elements in the array and M is the maximum length of string in array
 * <p>
 * Approach 2 using: Character count
 * T: O(N*M)
 * S: O(N*M) -> N is number of elements in the array and M is the maximum length of string in array
 */

public class Anagrams {
    public List<List<String>> groupAnagramBySort(List<String> input) {
        var map = new HashMap<String, List<String>>();
        for (String originalKey : input) {
            String sortedKey = getSortedKey(originalKey);
            var group = map.getOrDefault(sortedKey, new ArrayList<>());
            group.add(originalKey);
            map.put(sortedKey, group);
        }
        return new ArrayList<>(map.values());
    }

    public List<List<String>> groupAnagramByCharacterCount(List<String> input) {
        var map = new HashMap<String, List<String>>();
        var charInput = new int[26];
        for (String originalKey : input) {
            Arrays.fill(charInput, 0);
            var key =  getCharacterKey(originalKey, charInput);
            var group = map.getOrDefault(key, new ArrayList<>());
            group.add(originalKey);
            map.put(key, group);
        }
        return new ArrayList<>(map.values());
    }

    private String getSortedKey(String originalKey) {
        var chArray = originalKey.toCharArray();
        Arrays.sort(chArray);
        return new String(chArray);
    }

    private String getCharacterKey(String originalKey, int [] characterCount) {
        for (char ch : originalKey.toCharArray()) {
            characterCount[ch - 'a'] += 1;
        }

        var builder = new StringBuilder();
        for (int count: characterCount){
            builder.append("#");
            builder.append(count);
        }

        return builder.toString();
    }
}
