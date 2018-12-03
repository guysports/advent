package day1

import (
	"fmt"
	"strconv"
	"strings"
)

type (
	// Frequency - keep track of the frequency from the input
	Frequency struct {
		freq       int
		historic   []int
		repeatFreq int
	}
)

func validateInput(mod string) (int, error) {
	if !strings.HasPrefix(mod, "+") && !strings.HasPrefix(mod, "-") {
		return 0, fmt.Errorf("Invalid Input - Not a valid prefix")
	}

	i := 1
	if strings.HasPrefix(mod, "-") {
		i = 0
	}
	intmod, err := strconv.Atoi(mod[i:])

	if err != nil {
		return 0, fmt.Errorf("Invalid Input - Not a number")
	}

	return intmod, nil
}

// FindHistoric - Look for a repeat frequency when a new frequency number is added
func (f *Frequency) FindHistoric() bool {
	for _, oldHist := range f.historic {
		if oldHist == f.freq {
			f.repeatFreq = f.freq
			return true
		}
	}
	return false
}

// ProcessFrequencyInput - process the input string and determine frequency
func (f *Frequency) ProcessFrequencyInput(input string) {

	inputs := strings.Split(input, ",")
	found := false
	for {
		for _, input := range inputs {
			freqMod, err := validateInput(input)
			if err != nil {
				// Ignore invalid entry
				continue
			}
			f.freq += freqMod
			found = f.FindHistoric()
			f.historic = append(f.historic, f.freq)
			if found {
				return
			}
		}
	}
}

// GetFrequency - Show the calculated frequency
func (f *Frequency) GetFrequency() {
	fmt.Printf("Frequency is %d\n", f.freq)
	fmt.Printf("Repeat Frequency is %d\n", f.repeatFreq)
}
