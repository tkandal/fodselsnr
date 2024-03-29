// Package fodselsnr is a package that check Norwegian National Identity Numbers (NIN).  In addition
// to checking regular NINs, it should also be able to check so-called S-, D- and FS-numbers.  FS-number
// is a type of NIN that is used in the educational sector in Norway.
// NIN has the following format, it consists of 11 digit.  The first 6 digits are the birthdate
// with the format ddmmyy, and the last 5 are used for control and gender, on the format
// nngcc. nn is calculated from birthdate, g is gender, 0,2,4,6,8 for female and 1,3,5,7,9 for male,
// and lastly cc is the checksum for all the 9 proceeding digits.
// Ref. https://no.wikipedia.org/wiki/F%C3%B8dselsnummer (in Norwegian).
package fodselsnr

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

const (
	// IllegalControlSum is an illegal control sum.
	IllegalControlSum = 10
	// ZeroControlSum is the same as 0.
	ZeroControlSum = 11 // if the result becomes 11 - 0 => set to 0
	// FodselsnrLength is det legal length of NIN.
	FodselsnrLength = 11
	printFormat     = "%d-%02d-%02d"
	parseFormat     = "2006-01-02"
)

var (
	ninMatch *regexp.Regexp
)

func init() {
	ninMatch = regexp.MustCompile("^[\\d]{11}$")
}

// Check checks if a Norwegian National Identity Number (NIN) is legal.
// This function also checks if the NIN is a so-called legal S-, D- or FS-number.
// Returns true if the given NIN is either a regular NIN, an S-number, a D-number or an FS-number; false otherwise.
func Check(fnr string) bool {
	return Sjekk(fnr)
}

// isSNumber check if the given NIN is a so-called S-number.  A legal S-number
// has legal day (day > 0 and day < 32) and month + 50 (month > 50 and month < 63),
// year in the set [0-99], and the sum of the 7. and 8. digit > 9 and < 15
func isSNumber(fnr string) bool {
	day, month, year, err := parseDayMonthYear(fnr)
	if err != nil {
		return false
	}

	bornYear := calcYear(year)
	nr, err := strconv.Atoi(fnr[6:])
	if err != nil {
		return false
	}

	legal := (day > 0 && day < 32) && (month > 50 && month < 63) && (nr > 9000 && nr < 15000)
	if legal {
		legal = legal && parseIsLegal(bornYear, month-50, day)
	}
	return legal
}

// isDNumber check if the given NIN is a so-called D-number.  A legal D-number
// has day + 40 (day > 40 and day < 72) and legal month (month > 0 and month < 13),
// year in the set [0-99].
func isDNumber(fnr string) bool {
	day, month, year, err := parseDayMonthYear(fnr)
	if err != nil {
		return false
	}

	bornYear := calcYear(year)
	if !isCorrectIDNumber(bornYear, fnr) {
		return false
	}

	legal := (day > 40 && day < 72) && (month > 0 && month < 13)
	if legal {
		legal = legal && parseIsLegal(bornYear, month, day-40)
	}
	return legal
}

// isFSNumber check if the given NIN is a so-called FS-number. A legal FS-number
// has legal day (day > 0 and day < 32) and month + 50 (month > 50 and month < 63) and
// year in the set [0-99], and the sum of the last 5 digits > 89999.
func isFSNumber(fnr string) bool {
	day, month, year, err := parseDayMonthYear(fnr)
	if err != nil {
		return false
	}

	bornYear := calcYear(year)
	if !isCorrectIDNumber(bornYear, fnr) {
		return false
	}

	persNr, err := strconv.Atoi(fnr[6:])
	if err != nil {
		return false
	}
	legal := (day > 0 && day < 32) && (month > 50 && month < 63) && persNr > 89999
	if legal {
		legal = legal && parseIsLegal(bornYear, month-50, day)
	}
	return legal
}

// isRegular check if a given NIN is a regular NIN.  A regular NIN should
// have a legal day (day > 0 and day < 32) and a legal month (month > 0 and month < 13)
// and a year in the set [00 - 99].
func isRegular(fnr string) bool {
	day, month, year, err := parseDayMonthYear(fnr)
	if err != nil {
		return false
	}

	bornYear := calcYear(year)
	if !isCorrectIDNumber(bornYear, fnr) {
		return false
	}

	legal := (day > 0 && day < 32) && (month > 0 && month < 13)
	if legal {
		legal = legal && parseIsLegal(bornYear, month, day)
	}
	return legal
}

// Sjekk checks if a Norwegian National Identity Number (NIN) is legal.
// This function also checks if the NIN is a so-called legal S-, D- or FS-number.
// Returns true if the given NIN is either a regular NIN, an S-number, a D-number or an FS-number; false otherwise.
func Sjekk(fnr string) bool {
	if len(fnr) == 0 {
		return false
	}
	tmp := fnr
	if len(tmp) == (FodselsnrLength - 1) {
		tmp = "0" + tmp
	}
	if !ninMatch.Match([]byte(tmp)) {
		return false
	}

	if !isRegular(tmp) {
		if !isFSNumber(tmp) && !isDNumber(tmp) && !isSNumber(tmp) {
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

// 500–999 omfatter personer født i perioden 2000–2039.
// 000–499 omfatter personer født i perioden 1900–1999.
// 900–999 omfatter personer født i perioden 1940–1999.
// 500–749 omfatter personer født i perioden 1854–1899.
func isCorrectIDNumber(bornYear int, fnr string) bool {
	idNumber, err := strconv.Atoi(fnr[6:9])
	if err != nil {
		return false
	}

	if bornYear > 1999 && bornYear < 2040 {
		return idNumber >= 500 && idNumber <= 999
	}
	if bornYear > 1899 && bornYear < 2000 {
		return (idNumber >= 0 && idNumber <= 499) || (idNumber >= 900 && idNumber <= 999)
	}
	if bornYear > 1853 && bornYear < 1900 {
		return idNumber >= 500 && idNumber <= 749
	}
	return false
}

func parseDayMonthYear(fnr string) (int, int, int, error) {
	day, err := strconv.Atoi(fnr[0:2])
	if err != nil {
		return 0, 0, 0, err
	}
	month, err := strconv.Atoi(fnr[2:4])
	if err != nil {
		return 0, 0, 0, err
	}
	year, err := strconv.Atoi(fnr[4:6])
	if err != nil {
		return 0, 0, 0, err
	}
	return day, month, year, nil
}

func calcYear(year int) int {
	fnrYear := year + 2000
	if fnrYear > time.Now().Year() {
		fnrYear -= 100
	}
	return fnrYear
}

func parseIsLegal(year int, month int, day int) bool {
	if _, err := time.Parse(parseFormat, fmt.Sprintf(printFormat, year, month, day)); err != nil {
		return false
	}
	return true
}
