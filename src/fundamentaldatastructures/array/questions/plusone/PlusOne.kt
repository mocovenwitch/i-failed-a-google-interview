package fundamentaldatastructures.array.questions.plusone

import kotlin.math.pow

class PlusOne {

    /**
     * O(n) + O(n) = O(n)
     *
     * result: failed! the question doesn't mention the constrains!!
     */
    fun plusOne_1stTry(digits: IntArray): IntArray {
        // quickly handle empty
        if (digits.isEmpty()) return IntArray(0)

        // start
        var value = 0L
        var len = digits.size - 1
        digits.forEach {
            value += it * 10.toDouble().pow(len--).toLong()
        }

        value++

        len = value.toString().length
        val result = IntArray(len)
        var counter = 0
        var mod = 0L
        do {
            len--
            mod = (value % (10.toDouble().pow(len))).toLong()
            result[counter++] = (value / 10.toDouble().pow(len)).toInt()
            value = mod
        } while (mod > 0 && len > 0)

        return result
    }

    /**
     * Accepted answer. O(n)
     * - memory can be improved a bit, e.g. avoid to do a copy array
     */
    fun plusOne(digits: IntArray): IntArray {
        // empty array
        if (digits.isEmpty()) return IntArray(0)

        // start
        val len = digits.size

        var result = IntArray(len + 1)
        var carry = 1
        // count backwards
        for (i in len - 1 downTo 0) {
            var value = digits[i] + carry
            if (value >= 10) {
                carry = 1
                value %= 10
            } else {
                carry = 0
            }
            result[i+1] = value
        }
        // last carry
        result[0] = carry

        // trim
        if (result[0] == 0) {
            result = result.copyOfRange(1, result.size)
        }
        return result
    }
}
