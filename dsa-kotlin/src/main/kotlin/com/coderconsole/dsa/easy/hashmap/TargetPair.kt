package com.coderconsole.dsa.easy.hashmap

class TargetPair {
    fun isPairExistsByHashMap(input: List<Int>, target: Int): Boolean {
        val hashSet = HashSet<Any>()
        for (i in input) {
            if (hashSet.contains(i)) return true
            hashSet.add(target - i)
        }
        return false
    }

    fun isPairExistsByBruteForce(input: List<Int>, target: Int): Boolean {
        for (i in input.indices) {
            for (j in 0 until input.size - 1) {
                if (target - input[i] == input[j]) return true
            }
        }
        return false
    }
}