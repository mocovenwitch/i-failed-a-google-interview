package array.questions.plusone

fun main() {
    val case1= intArrayOf(1, 2, 3)
    val case2= intArrayOf(3)
    val case3= intArrayOf(0)

    val po = PlusOne()
    val result1 = po.plusOne(case1)
    val result2 = po.plusOne(case2)
    val result3 = po.plusOne(case3)

    result1.forEach { print("$it, ") }
    println()
    result2.forEach { print("$it, ") }
    println()
    result3.forEach { print("$it, ") }

}