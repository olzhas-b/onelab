package tetris

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"sort"
	"strings"
)

type Pair struct {
	Order string
	Sz    int
}
const (
	pattern = "struct.+\\{([^}]+)\\}"
)
func getDataTypeFromFile(path string) []string {
	file, _ := os.Open(path)
	defer file.Close()
	stat, _ := file.Stat()
	sz := stat.Size()
	bs := make([]byte, sz)

	file.Read(bs)
	file_s := string(bs)
	match, _ := regexp.MatchString(pattern, file_s)
	if match == false {
		panic("This file has not struct")
	}
	r1, _ := regexp.Compile(pattern)
	text_struct := r1.FindString(file_s)

	types := strings.Split(text_struct, "\n")
	types = types[1:len(types) - 1]
	return types
}

// точка входа
func GetTopMinimizedStruct(path string) []Pair {
	// get struct types
	elements := getDataTypeFromFile(path)
	// find all possible permutation
	permuatations := GetPermutations(len(elements))
	structsText := make([]string, 0, len(permuatations))
	ans := make([]Pair, 0, len(permuatations))
	for i := 0; i < len(permuatations); i++ {
		res := ""
		for j := 0; j < len(permuatations[i]); j++ {
			res += elements[permuatations[i][j]] + "\n"
		}
		structsText = append(structsText, res)
	}
	reCreateFile(structsText)
	calculateMemory(&ans, structsText)
	sort.Slice(ans, func(i, j int) bool {
		return ans[i].Sz < ans[j].Sz
	})
	return []Pair{ans[0], ans[len(ans) - 1]}
}
func reCreateFile(structsT []string) {
	file, _ := os.Create("tetris/generator/main.go")
	fileT := "package main\n\nimport (\n\t\"fmt\"\n\t\"unsafe\"\n)\n"
	for i := 0; i < len(structsT); i++ {
		s := fmt.Sprintf("type Changeable%d struct {\n %s \n}\n", i, structsT[i])
		fileT += s
	}
	fileT += "func main() {\n"
	for i := 0; i < len(structsT); i++ {
		s := fmt.Sprintf("\tfmt.Print(string(unsafe.Sizeof(Changeable%d{})))\n", i)
		fileT += s
	}
	fileT += "}"
	file.WriteString(fileT)
	file.Close()

}

func calculateMemory(ans *[]Pair, structsText []string)  {
	out, _ := exec.Command("/bin/sh", "-c", "go run tetris/generator/main.go").Output()
	output := string(out[:])
	for index, memoryS := range output {
		memoryInt := int(memoryS)
		k := Pair{structsText[index], memoryInt}
		*ans = append(*ans, k)
	}
}


func GetPermutations(sz int) [][]int {
	permutations := make([][]int, 0)
	storage := make([]int, sz)
	used := make([]bool, sz)
	recursion(0, storage, used, &permutations)
	return permutations
}
func recursion(index int, storage []int, used []bool, permutations *[][]int) {
	sz := len(used)
	if index == sz {
		storageCopy := make([]int, len(storage))
		copy(storageCopy, storage)
		*permutations = append(*permutations, storageCopy)
		return
	} else {
		for i := 0; i < sz; i++ {
			if  used[i] == false {
				used[i] = true
				storage[index] = i
				recursion(index + 1, storage, used, permutations)
				used[i] = false
			}
		}
	}
}