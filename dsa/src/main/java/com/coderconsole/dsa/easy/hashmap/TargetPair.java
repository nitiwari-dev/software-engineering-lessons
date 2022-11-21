package com.coderconsole.dsa.easy.hashmap;

import java.util.HashSet;
import java.util.List;

public class TargetPair {
    public boolean isPairExistsByHashMap(List<Integer> input, int target) {
        var hashSet = new HashSet<>();
        for (int i : input) {
            if (hashSet.contains(i)) return true;
            hashSet.add(target - i);
        }
        return false;
    }

    public boolean isPairExistsByBruteForce(List<Integer> input, int target) {
        for (int i = 0; i < input.size(); i++) {
            for (int j = 0; j < input.size() - 1; j++) {
                if (target - input.get(i) == input.get(j))
                    return true;
            }
        }

        return false;
    }
}
