package main

import (
	"log"
	"testing"
)

func TestMarkClaim(t *testing.T) {
	fab := fabric(5, 5)
	claims := make([]claim, 0)

	clm := claim{
		claimID:    1,
		cols:       3,
		rows:       2,
		leftMargin: 2,
		topMargin:  2,
	}
	claims = append(claims, clm)
	markClaim(fab, clm)

	clm = claim{
		claimID:    2,
		cols:       2,
		rows:       3,
		leftMargin: 1,
		topMargin:  2,
	}
	claims = append(claims, clm)
	markClaim(fab, clm)

	clm = claim{
		claimID:    3,
		cols:       3,
		rows:       3,
		leftMargin: 2,
		topMargin:  0,
	}
	claims = append(claims, clm)
	markClaim(fab, clm)

	clm = claim{
		claimID:    4,
		cols:       2,
		rows:       2,
		leftMargin: 0,
		topMargin:  0,
	}
	claims = append(claims, clm)
	markClaim(fab, clm)

	log.Println(overlapped(fab))
	log.Println(validateClaims(fab, claims))
	log.Println(fab)
}
