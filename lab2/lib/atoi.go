package lib

import (
	"errors"
)

var notNumber error = errors.New("FAILED convert into integer because it is not integer")

func Atoi(s string) (int, error) {
	if len(s) == 0 {
		return 1 << 31, notNumber
	}
	// recursion -> if our number less than we just calculate number without firstCharacter
	if string(s[0]) == "-" {
		recAns, err := Atoi(s[1:])
		return -1 * recAns, err
	}
	ans := 0
	for i := range s {
		if 48 > byte(s[i]) || byte(s[i]) > 58 {
			return 1 << 31, notNumber
		}
		ans += int(s[i]) - 48
		ans *= 10
	}
	return ans / 10, nil
}
