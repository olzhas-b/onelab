package lib

import (
	"log"
)

func RuneByIndex(s *string, i *int) (rune, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("panic recovered from function RuneByIndex(str, index)")
		}
	}()
	return rune((*s)[*i]), nil
}
