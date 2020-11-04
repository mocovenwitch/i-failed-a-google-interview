package questions

import (
	"math/rand"
	"time"
)

// ShuffleArray is
type ShuffleArray struct {
	Orignal []int
}

// Constructor is
func Constructor(nums []int) ShuffleArray {
	s := ShuffleArray{nums}
	return s
}

// Reset the array to its original configuration and return it.
func (s *ShuffleArray) Reset() []int {
	return s.Orignal
}

// Shuffle beat 15% in time
func (s *ShuffleArray) Shuffle() []int {
	shuffleed := make([]int, len(s.Orignal))
	copy(shuffleed, s.Orignal)
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(shuffleed), func(i, j int) { shuffleed[i], shuffleed[j] = shuffleed[j], shuffleed[i] })
	return shuffleed
}

// Shuffle1 beat 53.21% in time
func (s *ShuffleArray) Shuffle1() []int {
	tempArray := make([]int, len(s.Orignal))
	copy(tempArray, s.Orignal)
	for i := 0; i < len(tempArray); i++ {
		r := rand.Intn(i + 1)
		tempArray[i], tempArray[r] = tempArray[r], tempArray[i]
	}
	return tempArray
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
