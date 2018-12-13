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
	claims := make([]claim, 0)

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

		markClaim(fab, clm)
		claims = append(claims, clm)
	}
	fmt.Println(overlapped(fab))
	fmt.Println(validateClaims(fab, claims))
}

func fabric(x, y int) [][]int {
	fab := make([][]int, x)
	for i := range fab {
		fab[i] = make([]int, y)
	}

	return fab
}

func overlapped(fab [][]int) int {
	overlapped := 0
	for _, row := range fab {
		for j := range row {
			if row[j] > 1 {
				overlapped++
			}
		}
	}
	return overlapped
}

func markClaim(fab [][]int, clm claim) {
	for i := clm.leftMargin; i < clm.leftMargin+clm.cols; i++ {
		for j := clm.topMargin; j < clm.topMargin+clm.rows; j++ {
			fab[i][j]++
		}
	}
}

func validateClaims(fab [][]int, claims []claim) int {
	for _, claim := range claims {
		validClaim := true
		for i := claim.leftMargin; i < claim.leftMargin+claim.cols; i++ {
			for j := claim.topMargin; j < claim.topMargin+claim.rows; j++ {
				if fab[i][j] > 1 {
					validClaim = false
				}
			}
		}
		if validClaim {
			return claim.claimID
		}
	}
	return -1
}
