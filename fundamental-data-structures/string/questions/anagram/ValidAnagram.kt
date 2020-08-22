package string.questions.anagram

class ValidAnagram {
    // O(n)
    fun isAnagram(s: String, t: String): Boolean {
        val map: HashMap<Char, Int> = hashMapOf()
        for (a in s.toCharArray()) {
            map[a] = ((map[a] ?: 0) + 1)
        }
        for (b in t.toCharArray()) {
            // if can't find in s, invalid anagram
            map[b] ?: run { return false }

            // continue
            map[b] = map[b]!! - 1
        }

        map.forEach {
            if (it.value != 0) {
                return false
            }
        }
        return true

    }
}
