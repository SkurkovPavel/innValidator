package innValidator

import (
	"regexp"
)

// Validate method of checking the control number for an individual or legal entity
func Validate(inn string) bool {

	if len(inn) == 12 {
		return validTinNumber(inn, 2, 1) && validTinNumber(inn, 1, 0)
	} else {
	    // test
		return validTinNumber(inn, 1, 2)
	}

}

func validTinNumber(inn string, offset, arrOffset int) bool {

	const innRegexp = `^\d{10}$|^\d{12}$`

	matched, err := regexp.Match(innRegexp, []byte(inn))
	if err != nil {
		return false
	}

	if !matched {
		return false
	}

	checkArr := []uint8{3, 7, 2, 4, 10, 3, 5, 9, 4, 6, 8}
	var sum uint8 = 0
	var length = len(inn)

	for i := 0; i != (length - offset); i++ {
		sum = sum + ((inn[i] - '0') * checkArr[i+arrOffset])

	}
	res := sum%11%10 == inn[length-offset]-'0'

	return res

}

