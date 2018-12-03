package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/guysports/advent/pkg/day1"
	"github.com/guysports/advent/pkg/day2"
)

func main_day1() {
	input := os.Args[1]
	frequency := day1.Frequency{}
	frequency.ProcessFrequencyInput(input)
	frequency.GetFrequency()
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {
	inputFile := os.Args[1]
	keys, _ := readLines(inputFile)

	checksum := day2.CreateCheckSum(keys)
	fmt.Printf("Checksum is %d\n", checksum)
}
