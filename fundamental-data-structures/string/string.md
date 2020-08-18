
# String

## In Java


## Golang

Question:`How to use two pointers in golang for-loop?`

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