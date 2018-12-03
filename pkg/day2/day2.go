package day2

import (
	"fmt"
	"strings"
)

type (
	// Occurences - keeps track of the occurences of letters in a key
	Occurences struct {
		two   int
		three int
	}
)

// ProcessKey - returns the number of 2 and 3 letter occurences in a key
func ProcessKey(key string) *Occurences {

	occurence := &Occurences{}

	for _, char := range key {
		count := strings.Count(key, fmt.Sprintf("%c", char))
		if count == 2 {
			occurence.two = 1
		} else if count > 2 {
			occurence.three = 1
		}
		if occurence.two == 1 && occurence.three == 1 {
			break
		}
	}

	return occurence
}

// CreateChecksum - calculate a checksum from the keys
func CreateChecksum(keys []string) int {
	twos := 0
	threes := 0
	for _, key := range keys {
		occurence := ProcessKey(key)
		if occurence.two == 1 {
			twos++
		}
		if occurence.three == 1 {
			threes++
		}
	}
	checksum := twos * threes

	return checksum
}
