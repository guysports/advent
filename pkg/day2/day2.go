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

// FindBoxID - Find the strings that differ by a single character
func FindBoxID(keys []string) string {
	for pos, key := range keys {
		// Test keys after key
		for i := pos + 1; i < len(keys); i++ {
			// Check key against laterkey
			count := 0
			different := 0
			for j := 0; j < len(key); j++ {
				if key[j] != keys[i][j] {
					count++
					different = j
				}
			}
			if count == 1 {
				return key[(different-1):] + key[:different]
			}
		}
	}
	return "nomatch"
}
