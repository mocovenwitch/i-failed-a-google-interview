package fundamentaldatastructures.array

fun main() {
    normalArray()
    primitiveArray()
}

fun primitiveArray() {
    // An array of ints. When targeting the JVM, instances of this class are represented as `int[]`.
    val intArray = IntArray(5)
    print(intArray)
}

fun normalArray() {
    val array = arrayOf(1, 2)
    array.forEach {
        println(it)
    }
}