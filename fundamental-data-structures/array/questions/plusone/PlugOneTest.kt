package array.questions.plusone

fun main() {
    val case1= intArrayOf(1, 2, 3)
    val case2= intArrayOf(3)
    val case3= intArrayOf(0)
    val case4 = intArrayOf(9,8,7,6,5,4,3,2,1,0)
    val case5 = intArrayOf(6,1,4,5,3,9,0,1,9,5,1,8,6,7,0,5,5,4,3)
    val case6= intArrayOf(9,9)


    val po = PlusOne()
    val result1 = po.plusOne(case1)
    val result2 = po.plusOne(case2)
    val result3 = po.plusOne(case3)
    val result4 = po.plusOne(case4)
    val result5 = po.plusOne(case5)
    val result6 = po.plusOne(case6)

    result1.forEach { print("$it, ") }
    println()
    result2.forEach { print("$it, ") }
    println()
    result3.forEach { print("$it, ") }
    println()
    result4.forEach { print("$it, ") }
    println()
    result5.forEach { print("$it, ") }
    println()
    result6.forEach { print("$it, ") }

}