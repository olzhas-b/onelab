package lib

func Reverse(s string) string {
	ns := ""
	var storage []rune
	for _, value := range s {
		storage = append(storage, value)
	}
	for i := len(storage) - 1; i >= 0; i-- {
		ns += string(storage[i])
	}
	return ns
}