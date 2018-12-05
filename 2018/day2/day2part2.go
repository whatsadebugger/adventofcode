package main

import (
	"bufio"
	"fmt"
	"os"
)

// part 2 involves finding the two boxes that differ by only one character
// and then finding the letters common between those two boxes. There will be
// only one correct box.
func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	_ = scanner

	boxes := make([]string, 0) // 250 is the number of boxes in out input.txt

	for scanner.Scan() {
		boxes = append(boxes, scanner.Text())
	}

	for i := 0; i < len(boxes); i++ {
		box1 := boxes[i]
		for j := 0; j < len(boxes); j++ {
			if j == i {
				continue
			}
			box2 := boxes[j]
			if IsOneCharApart(box1, box2) {
				fmt.Println("Found a match with ", box2)
				fmt.Println("Matching box characters: ", CommonCharacters(box1, box2))
				break
			}
		}
	}
}

func IsOneCharApart(box, anotherBox string) bool {
	IsOneCharApart := false
	count := 0
	for i := range box {
		if box[i] != anotherBox[i] {
			count++
		}
	}
	if count == 1 {
		IsOneCharApart = true
	}
	return IsOneCharApart
}

func CommonCharacters(box, another string) string {
	common := make([]byte, 0)

	for i := range box {
		if box[i] == another[i] {
			common = append(common, box[i])
		}
	}

	return string(common)
}
