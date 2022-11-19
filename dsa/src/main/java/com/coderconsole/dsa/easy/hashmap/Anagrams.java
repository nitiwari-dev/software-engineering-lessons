package com.coderconsole.dsa.easy.hashmap;


import java.util.ArrayList;
import java.util.List;

/**
 * Given an array of strings strs, group the anagrams together. You can return the answer in any order.

 An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

 Example 1:

 Input: strs = ["eat","tea","tan","ate","nat","bat"]
 Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
 Example 2:

 Input: strs = [""]
 Output: [[""]]
 Example 3:

 Input: strs = ["a"]
 Output: [["a"]]


 Approach 1 using: Sorting
 T: O(N*MLogM)
 S: O(N*M) -> N is number of elements in the array and M is the maximum length of string in array

 Approach 2 using: Character count
 T: O(N*M)
 S: O(N*M) -> N is number of elements in the array and M is the maximum length of string in array
*/

public class Anagrams {
    public List<List<String>> groupAnagramTogetherWithSorting(List<String> input) {
        return List.of(input);
    }
}
