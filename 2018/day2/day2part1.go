package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	_ = scanner

	two := 0
	three := 0

	for scanner.Scan() {
		ID := scanner.Text()
		addedTwo := false
		addedThree := false

		seen := make(map[string]int)
		for _, c := range ID {
			seen[string(c)]++
		}
		for _, v := range seen {
			if v == 2 && !addedTwo {
				two++
				addedTwo = !addedTwo
			} else if v == 3 && !addedThree {
				three++
				addedThree = !addedThree
			}
		}
	}

	fmt.Println("Checksum:", two*three)
}
