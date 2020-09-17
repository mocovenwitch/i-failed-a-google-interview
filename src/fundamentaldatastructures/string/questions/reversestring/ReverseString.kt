package fundamentaldatastructures.string.questions.reversestring

object ReverseString {

    @JvmStatic
    fun reverseString(s: CharArray): Unit {
        val len = s.size
        var r = len -1
        var l = 0
        while (l < r) {
            val t = s[l]
            s[l] = s[r]
            s[r] = t
            r--
            l++
        }
    }
}