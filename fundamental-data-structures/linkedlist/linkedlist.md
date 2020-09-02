
# LinkedList

## Java

- `List` has a `specified order` and provides `indexed access` to them
- `ArrayList` is an implementation of List, has order for sure, based on the index.
- `LinkedList` is a doubly-linked List, and is an implementation of List as well, ordered.
    - it has a Node class managing the order, which has previous and next reference, but invisible to outside

```
private static class Node<E> {
    E item;
    Node<E> next;
    Node<E> prev;

    Node(Node<E> prev, E element, Node<E> next) {
        this.item = element;
        this.next = next;
        this.prev = prev;
    }
}
```
## Kotlin

- listOf(), immutable, when it's empty, it is an EmptyList, which is an implementation of List
- listOf(1, 2, 3), when it has elements, it is an ArrayList, but `immutable`!
- arrayListOf(), it's simply an Java ArrayList, mutable!

[Code example](MyLinkedList.kt)

## Go

Ha, Golang does not have `while` keyword, but user for instead:

```
for head.Next != nil {
    n = head.Next
    head.Next = p
    p = head
    head = n
}
```