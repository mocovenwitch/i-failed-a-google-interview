package string

fun main() {
    share()
    notShare()
}

private fun share() {
    val a = "share"
    val b = "share"
    var c = ""
    c = "share"

    println(a === b)
    println(a === c)
}

private fun notShare() {
    val na = String().plus("not-share")
    val nb = String().plus("not-share")
    println(nb === na)
}