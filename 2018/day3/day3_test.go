package main

import (
	"log"
	"testing"
)

func TestMarkClaim(t *testing.T) {
	fab := fabric(5, 5)
	clm := claim{
		claimID:    1,
		rows:       2,
		cols:       3,
		leftMargin: 2,
		topMargin:  2,
	}
	markClaim(fab, clm)
	clm = claim{
		claimID:    2,
		rows:       3,
		cols:       2,
		leftMargin: 1,
		topMargin:  2,
	}

	markClaim(fab, clm)
	log.Println(overlapped(fab))
	log.Println(fab)

}
