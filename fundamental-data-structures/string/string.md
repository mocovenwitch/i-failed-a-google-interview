
# String

## In Java

Whenever a string is constructed using a loop that iterates more than just a few times, it is usually better to create a StringBuffer or StringBuilder object and append to that. 
Because such buffers are based on mutable character arrays, which do not require a new string to be created for each concatenation, they can reduce the cost of repeatedly growing the string.

To choose between `StringBuffer` and StringBuilder, check if the new buffer object may be accessed by different threads while in use. If `multi-thread safety` is required, use a StringBuffer. For purely local string buffers, you can avoid the overhead of synchronization by using a StringBuilder.

- StringBuffer for multi-thread safety
- StringBuilder for normal usage

### fun fact
When we create a String object using the new() operator, it always creates a new object in heap memory. On the other hand, if we create an object using String literal syntax e.g. “Baeldung”, it may return an existing object from the String pool, if it already exists.
Before Java 7, String pool is at PermGen, can't be garbage colleced. While after Java 7, String pool is in the heap too.

- new String("I am in the heap")
- String s = "I am in the stack String pool, happy to share"

[String code example](JavaString.java)

## Kotlin
When Kotlin runs in JVM, similar to Java

- String() create a brand-new object in the heap
- val a = "hello" would share the same String previously in the String pool

[Kotlin code example](KotlinString.kt)

## Golang
string is the set of all strings of 8-bit bytes, conventionally but not necessarily representing UTF-8-encoded text. A string may be empty, but not nil. Values of string type are immutable. It does not look like that golang has a String pool for sharing. [mark as a question]

Try this `go tool compile "-m" main.go`

[Golang String](main.go)


### Questions

How to use two pointers in golang for-loop?

One pointer is defined ahead of for-loop, is this the normal way to do it?
```
r := size - 1
if size > 1 {
    for l := 0; l < r; l++ {
        t := s[l]
        s[l] = s[r]
        s[r] = t
        r--
    }
}
```

### Summary
We know this for sure, but still make mistake in practice.

Even O(n) == O(n) + O(n) + O(n), still try to reduce it.

- Unfixed array slower than fixed array
- map is more expensive than primitive array