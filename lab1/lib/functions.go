package lib

import (
	"errors"
	"github.com/gofrs/uuid"
	"math"
	"strconv"
)

func LowerToUpper(word string) string {
	b := []byte(word)
	for i := range b {
		if 97 <= b[i] && b[i] <= 122 {
			b[i] -= 32
		}
	}
	return string(b)
}
func UpperToLower(word string) string {
	b := []byte(word)
	for i := range b {
		if 65 <= b[i] && b[i] <= 90 {
			b[i] += 32
		}
	}
	return string(b)
}

func FindRootQuEq(a, b, c float64) (float64, float64, error) {
	d := math.Sqrt((b * b) - (4 * a * c))
	if d == 0 {
		x := -b / (2 * a)
		return x, x, nil
	} else if d > 0 {
		x1 := (-b + d) / (2 * a)
		x2 := (-b - d) / (2 * a)
		return x1, x2, nil
	} else {
		return 0, 0, errors.New("dont have root")
	}
}


func GetId() string {
	u2, err := uuid.NewV4()
	if err != nil {
		print(err)
	}
	id := ""
	for _, b := range u2 {
		id += strconv.Itoa(int(b))
	}
	return id
}

