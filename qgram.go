package edlib

import (
	"math"
)

// QgramSimilarity compute the q-gram similarity between two strings
// Takes two strings as parameters, a split length which defines the k-gram shingle length
func QgramSimilarity(str1, str2 string, splitLength int) float32 {
	splittedStr1 := Shingle(str1, splitLength)
	splittedStr2 := Shingle(str2, splitLength)

	union := make(map[string]int)
	for i := range splittedStr1 {
		union[i] = 0
	}
	for i := range splittedStr2 {
		union[i] = 0
	}

	res := float32(0)

	for i := range union {
		res += float32(math.Abs(float64(splittedStr1[i] - splittedStr2[i])))
	}

	return res
}
