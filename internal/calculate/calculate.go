package calculate

import (
	"strings"
)

type CharacterCounts map[string]int

type KeyVal struct {
	Key string `json:"word"`
	Val int    `json:"score"`
}

type Alphabet struct {
	Letters []KeyVal
}

type FiveLetters struct {
	Word []Alphabet
}

const lut = "abcdefghijklmnopqrstuvwxyz"

func getIndexOfRune(r rune) int {
	return strings.IndexRune(lut, r)
}

func newAlphabet() *Alphabet {
	var a Alphabet
	for _, v := range lut {
		a.Letters = append(a.Letters, KeyVal{
			Key: string(v),
			Val: 0,
		})
	}
	return &a
}

func newFiveLetters() *FiveLetters {
	var fl FiveLetters

	for x := 0; x < 5; x++ {
		fl.Word = append(fl.Word, *newAlphabet())
	}

	return &fl
}

func GetCharacterCounts(s []string) *FiveLetters {
	fl := newFiveLetters()

	for _, word := range s {
		for i, char := range word {
			fl.Word[i].Letters[getIndexOfRune(char)].Val += 1
		}
	}
	return fl
}

func swap(kv []KeyVal, x int, y int) {
	tmp := kv[x]
	kv[x] = kv[y]
	kv[y] = tmp
}

func sortKVByValue(kv []KeyVal) []KeyVal {
	var minIdx int
	for i := range kv {
		minIdx = i
		for j := i + 1; j < len(kv); j++ {
			if kv[j].Val < kv[minIdx].Val {
				minIdx = j
			}
			swap(kv, minIdx, i)
		}
	}

	var reverseKV []KeyVal
	for x := len(kv) - 1; x != 0; x-- {
		reverseKV = append(reverseKV, kv[x])
	}
	return reverseKV
}

func GeneratePowerStarters(list []string, fl FiveLetters, n int) []KeyVal {
	if n < 0 || n > 26 {
		n = 26
	}

	for i, word := range fl.Word {
		fl.Word[i] = Alphabet{
			Letters: sortKVByValue(word.Letters),
		}
	}

	var (
		count        int
		intermediate []KeyVal
	)

	for a, one := range fl.Word[0].Letters[:n] {
		for b, two := range fl.Word[1].Letters[:n] {
			for c, three := range fl.Word[2].Letters[:n] {
				for d, four := range fl.Word[3].Letters[:n] {
					for e, five := range fl.Word[4].Letters[:n] {
						word := one.Key + two.Key + three.Key + four.Key + five.Key
						sum := a + b + c + d + e
						for _, val := range list {
							if word == val && count < n && IsUniqueCharacters(word) {
								intermediate = append(intermediate, KeyVal{
									Key: word,
									Val: 130 - sum,
								})
							}
						}
					}
				}
			}
		}
	}

	return sortKVByValue(intermediate)
}

func IsUniqueCharacters(word string) bool {
	wordMap := make(map[rune]bool)
	for _, char := range word {
		if !wordMap[char] {
			wordMap[char] = true
		} else {
			return false
		}
	}
	return true
}
