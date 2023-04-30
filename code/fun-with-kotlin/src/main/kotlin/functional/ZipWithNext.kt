package functional

import kotlin.math.max

/**
 * Given a list of number find the max adjacent product and return the same. Assume there will always be number > 2
 */

class ZipWithNext {

    fun legacyLargestProduct(list: MutableList<Int>): Int {
        var i = 0
        var j = 1
        var maxAdjacentProduct = list.product(i++, j++)

        while (j < list.size){
            maxAdjacentProduct = max(maxAdjacentProduct, list.product(i++, j++))
        }
        return maxAdjacentProduct
    }

    fun functionalLargestProduct(input: MutableList<Int>): Int {
        return input.zipWithNext(Int::times).max()
    }

    private fun List<Int>.product(i: Int, j : Int): Int {
        return this[i] * this[j]
    }
}
