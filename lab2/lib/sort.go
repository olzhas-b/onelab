package lib

import (
	"errors"
	"os"
	"regexp"
	"sort"
	"strings"
)
const (
	pattern = "import.+\\(([^)]+)\\)"
	patternForLib = "\"([^\"]*)\""
)
var impossible error = errors.New("Failed to sort imports in the golang")

func SortImports(path string) error  {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return impossible
	}
	stat, _ := file.Stat()
	sz := stat.Size()
	bs := make([]byte, sz)

	file.Read(bs)
	text := string(bs)
	match, _ := regexp.MatchString(pattern, text)
	if match == false {
		return nil
	}

	// get import (...)
	r, _ := regexp.Compile(pattern)
	curr_text := r.FindString(text)
	println(curr_text)
	// get "github"
	r2 , _ := regexp.Compile(patternForLib)
	t := r2.FindAllString(curr_text, -1)

	sort.Strings(t)
	new_text := "import (\n"
	for _, library := range t {
		new_text += "\t" + library + "\n"
	}
	new_text += ")"

	// replace previous import to new
	text = strings.Replace(text, curr_text, new_text, 1)
	file, err = os.Create(path)
	defer file.Close()
	if err != nil {
		return err
	}
	file.WriteString(text)
	return nil
}
