package utils

import (
	"math"
	"regexp"
	"strconv"
)

func CalcTo(w float64, f string, t string) float64 {
	sr := map[string]float64{
		"Ki": 1,
		"Mi": 2,
		"Gi": 3,
	}

	ds := sr[f] - sr[t]
	if ds > 0 {
		return w * math.Pow(1024, ds)
	} else if ds < 0 {
		return w / math.Pow(1024, math.Abs(ds))
	} else {
		return w
	}
}

func SplitNumberUnit(s string) (int, string) {
	reNum := regexp.MustCompile(`\d+`)
	reUnit := regexp.MustCompile(`[^\d]+`)

	numStr := reNum.FindString(s)
	unit := reUnit.FindString(s)

	num, _ := strconv.Atoi(numStr)
	return num, unit
}
