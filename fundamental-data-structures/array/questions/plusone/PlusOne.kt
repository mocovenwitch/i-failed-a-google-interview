package array.questions.plusone

import kotlin.math.pow

class PlusOne {

    /**
     * O(n) + O(n) = O(n)
     */
    fun plusOne(digits: IntArray): IntArray {
        // quickly handle empty
        if (digits.isEmpty()) return IntArray(0)

        // start
        var value = 0
        var len = digits.size-1
        digits.forEach {
            value += it * 10.toDouble().pow(len--).toInt()
        }

        value++

        len = value.toString().length
        val result = IntArray(len)
        var counter = 0
        var mod = 0
        do {
            len--
            mod = (value % (10.toDouble().pow(len))).toInt()
            result[counter++] = (value / 10.toDouble().pow(len)).toInt()
            value = mod
        } while (mod > 0 && len >0)

        return result
    }
}
