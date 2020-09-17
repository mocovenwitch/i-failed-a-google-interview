package fundamentaldatastructures.linkedlist

import java.util.*

fun main() {
    // here, list is actually an ArrayList (found out by reading the source code), but immutable
    val list = listOf(1, 2, 3)
    // list.add () immutable

    // here, equals Java ArrayList, mutable
    val al = arrayListOf(1, 2, 3)
    al.add(4)
    al.forEach { print(it) }

    // Now, s is a Doubly-linked list.
    val s = LinkedList(list)
    s.forEach {
        print(it)
    }
}