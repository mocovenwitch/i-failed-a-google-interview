package blockchain

import java.security.MessageDigest

fun hashMessage(zeros: Int, value: String): Pair<String, String> {
    val zeroString = zeros.getZeros()

    (0..100000).forEach { index ->
        val result = (value+index).sha256()
        // println(result)
        if (result.subSequence(0, zeros) == zeroString) {
            return Pair((value+index), result)
        }
    }

    return Pair("", "")
}

private fun Int.getZeros(): String {
    var zeroString = ""
    (1..this).forEach{ _ ->
        zeroString += "0"
    }
    return zeroString
}

private fun String.sha256(): String {
    return hashString(this, "SHA-256")
}

private fun hashString(input: String, algorithm: String): String {
    return MessageDigest
            .getInstance(algorithm)
            .digest(input.toByteArray())
            .fold("", { str, it -> str + "%02x".format(it) })
}

/**
 * To find the min number for hashing a message that get X zeros in the front
 */
fun main() {
//    println("found: " + blockchain.hashMessage(3, "Bitcoin is cool "))

    val message = "0038CD488F8D215901F1CD6B5280E7590C62E2B06FA9429E92871C6B9AF641C0"
    println("found: " + hashMessage(3, message))
}