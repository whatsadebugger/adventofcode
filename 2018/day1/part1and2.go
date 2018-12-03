package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Part1 was to calculate the frequency after 1 pass of the input which was done easily
// in the first iteration of the code.
// Part2 was to iterate through the file until a duplicate frequency has been found
func main() {
	total := 0
	// seen is a map of ints that have been seen
	seen := make(map[int]struct{})
	// seenTwiceFirst will not be empty as soon as we find our duplicate frequency
	seenTwiceFirst := ""

	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	for seenTwiceFirst == "" {
		scanner := bufio.NewScanner(f)
		f.Seek(0, 0)

		for scanner.Scan() {
			txt := scanner.Text()
			A, _ := strconv.Atoi(txt)
			if _, ok := seen[total]; ok {
				seenTwiceFirst = strconv.Itoa(total)
				break
			}
			seen[total] = struct{}{}
			total += A
		}
	}
	fmt.Println("seenTwiceFirst", seenTwiceFirst)
}
