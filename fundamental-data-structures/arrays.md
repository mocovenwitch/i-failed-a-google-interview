
# Arrays

## Java
The facts that need to keep in mind:
- int[], or long[] are an array of primitive types, it is an object, so the elements are stored in `heap` other than stack.
- The size of each element of Array is the same. e.g. an int[], each element is 4 bytes. An Object[], each element is also the same, because the value of each element is NOT the object itself, it is the reference of object, they are always the same.
- Array is a fixed-size collection of items, same in Golang.
- Arrays.copyOf() actually call System.arraycopy() underneath, and then call native library to finish the operation

## Golang
- Golang array is a fixed-size collection of items of the same type
- Where do the values sit?
- 
