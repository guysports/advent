package main

import (
	"os"

	"github.com/guysports/advent/pkg/day1"
)

func main() {
	input := os.Args[1]
	frequency := day1.Frequency{}
	frequency.ProcessFrequencyInput(input)
	frequency.GetFrequency()
}
