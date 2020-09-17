package fundamentaldatastructures.string.questions.anagram

import fundamentaldatastructures.string.questions.anagram.ValidAnagram

fun main() {

    val case1a = "anagram"
    val case1b = "nagaram"
    println(ValidAnagram().isAnagram(case1a, case1b))


    val case2a = "rat"
    val case2b = "car"
    println(ValidAnagram().isAnagram(case2a, case2b))

    val case3a = "anagram"
    val case3b = "nagaramm"
    println(ValidAnagram().isAnagram(case3a, case3b))
}