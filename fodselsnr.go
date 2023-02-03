package fodselsnr

import (
	"strconv"
)

const (
	IllegalControlSum = 10
	ZeroControlSum    = 11 // if the result becomes 11 - 0 => set to 0
	FodselsnrLength   = 11
)

// Check checks if a norwegian national identity number (NIN) is legal.
// This function does not check if D-, H-, FH- or S-numbers are legal.
func Check(fnr string) bool {
	return Sjekk(fnr)
}

// isSnumber check if the given NIN is a so-called S-number.  A legal S-number
// has legal day (day > 0 and day < 32) and month + 50 (month > 50 and month < 63)
// and the sum of the 7. and 8. digit >= 10 and <= 14
func isSnumber(fnr string) bool {
	day, err := strconv.Atoi(fnr[0:2])
	if err != nil {
		return false
	}
	month, err := strconv.Atoi(fnr[2:4])
	if err != nil {
		return false
	}
	nr, err := strconv.Atoi(fnr[6:8])
	if err != nil {
		return false
	}
	return (day > 0 && day < 32) && (month > 50 && month < 63) && (nr >= 10 && nr <= 14)
}

// isDnumber check if the given NIN is a so-called D-number.  A legal D-number
// has day + 40 (day > 40 and day < 72) and legal month (month > 0 and month < 13)
// and the 7. digit == 0
func isDnumber(fnr string) bool {
	day, err := strconv.Atoi(fnr[0:2])
	if err != nil {
		return false
	}
	month, err := strconv.Atoi(fnr[2:4])
	if err != nil {
		return false
	}
	nr, err := strconv.Atoi(fnr[6:7])
	if err != nil {
		return false
	}
	return (day > 40 && day < 72) && (month > 0 && month < 13) && nr == 0
}

// isFSnumber check if the given NIN is a so-called FS-number. A legal FS-number
// has legal day (day > 0 and day < 32) and month + 50 (month > 50 and month < 63)
// and the sum of the last 5 digit >= 90000
func isFSnumber(fnr string) bool {
	day, err := strconv.Atoi(fnr[0:2])
	if err != nil {
		return false
	}
	month, err := strconv.Atoi(fnr[2:4])
	if err != nil {
		return false
	}
	persNr, err := strconv.Atoi(fnr[6:])
	if err != nil {
		return false
	}
	return day < 32 && (month > 50 && month < 63) && persNr >= 90000
}

func isRegular(fnr string) bool {
	day, err := strconv.Atoi(fnr[0:2])
	if err != nil {
		return false
	}
	month, err := strconv.Atoi(fnr[2:4])
	if err != nil {
		return false
	}
	return day < 32 && month < 13
}

// 075863 00000

// Sjekk checks if a norwegian national identity number (NIN) is legal.
// This function also checks if the NIN is a so-called legal S-, D- or FS-number
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

	if !isRegular(tmp) {
		if !isSnumber(tmp) && !isDnumber(tmp) && !isFSnumber(tmp) {
			return false
		}
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
