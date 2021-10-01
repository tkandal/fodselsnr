package fodselsnr

import (
	"strconv"
)

const (
	IllegalControlSum = 10
	ZeroControlSum    = 11 // if the result becomes 11 - 0 => set to 0
	FodselsnrLength   = 11
)

func Check(fnr string) bool {
	return Sjekk(fnr)
}

// Sjekk checks if a norwegian national identity number (NIN) is legal.
// This function does not check if D-, H-, FH- or S-numbers are legal.
func Sjekk(fnr string) bool {
	if len(fnr) == 0 {
		return false
	}
	tmp := fnr
	if len(tmp) == (FodselsnrLength - 1) {
		tmp = "0" + tmp
	}
	if len(tmp) != FodselsnrLength {
		return false
	}

	day1, err := strconv.Atoi(tmp[0:1])
	if err != nil {
		return false
	}
	day2, err := strconv.Atoi(tmp[1:2])
	if err != nil {
		return false
	}
	month1, err := strconv.Atoi(tmp[2:3])
	if err != nil {
		return false
	}
	month2, err := strconv.Atoi(tmp[3:4])
	if err != nil {
		return false
	}
	aar1, err := strconv.Atoi(tmp[4:5])
	if err != nil {
		return false
	}
	aar2, err := strconv.Atoi(tmp[5:6])
	if err != nil {
		return false
	}
	i1, err := strconv.Atoi(tmp[6:7])
	if err != nil {
		return false
	}
	i2, err := strconv.Atoi(tmp[7:8])
	if err != nil {
		return false
	}
	i3, err := strconv.Atoi(tmp[8:9])
	if err != nil {
		return false
	}
	kontroll1, err := strconv.Atoi(tmp[9:10])
	if err != nil {
		return false
	}
	kontroll2, err := strconv.Atoi(tmp[10:])
	if err != nil {
		return false
	}

	kalk1 := 11 - ((3*day1 + 7*day2 + 6*month1 + month2 + 8*aar1 + 9*aar2 + 4*i1 + 5*i2 + 2*i3) % 11)
	if kalk1 == IllegalControlSum {
		return false
	}
	if kalk1 == ZeroControlSum {
		kalk1 = 0
	}

	kalk2 := 11 - ((5*day1 + 4*day2 + 3*month1 + 2*month2 + 7*aar1 + 6*aar2 + 5*i1 + 4*i2 + 3*i3 + 2*kalk1) % 11)
	if kalk2 == IllegalControlSum {
		return false
	}
	if kalk2 == ZeroControlSum {
		kalk2 = 0
	}
	return kontroll1 == kalk1 && kontroll2 == kalk2
}
