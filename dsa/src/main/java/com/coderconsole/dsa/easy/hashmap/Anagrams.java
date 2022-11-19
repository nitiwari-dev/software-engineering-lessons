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
    public List<List<String>> groupAnagramTogetherWithSorting(List<String> input) {

        var map = new HashMap<String, List<String>>();
        for (String originalKey : input) {
            String sortedKey = getSortedKey(originalKey);
            if (map.containsKey(sortedKey)) {
                var group = map.get(sortedKey);
                group.add(originalKey);
                map.put(sortedKey, group);
            } else {
                var group = new ArrayList<String>();
                group.add(originalKey);
                map.put(sortedKey, group);
            }
        }
        return new ArrayList<>(map.values());
    }

    private String getSortedKey(String originalKey) {
        var chArray = originalKey.toCharArray();
        Arrays.sort(chArray);
        return new String(chArray);
    }
}
