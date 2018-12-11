package main

import (
	"bufio"
	"fmt"
	"os"
)

type claim struct {
	claimID    int
	cols       int
	rows       int
	leftMargin int
	topMargin  int
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	fab := fabric(1000, 1000)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		clm := claim{}

		fmt.Sscanf(
			line,
			"#%d @ %d,%d: %dx%d",
			&clm.claimID,
			&clm.leftMargin,
			&clm.topMargin,
			&clm.cols,
			&clm.rows,
		)

		// mark claims in a way that we can detect overlaps?
		markClaim(fab, clm)
		// find overlaps

	}
}

func fabric(x, y int) [][]int {
	fab := make([][]int, x)
	for i := range fab {
		fab[i] = make([]int, y)
	}

	return fab
}

func markClaim(fab [][]int, clm claim) {

}
