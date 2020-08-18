package string.questions.reversestring

fun main() {
    val case1 = charArrayOf('h','e','l','l','o')
    case1.forEach { print("$it,") }
    println()
    ReverseString.reverseString(case1)
    case1.forEach { print("$it,") }

    val case2 = charArrayOf('H','a','n','n','a','h')
    println()
    case2.forEach { print("$it,") }
    println()
    ReverseString.reverseString(case2)
    case2.forEach { print("$it,") }
}