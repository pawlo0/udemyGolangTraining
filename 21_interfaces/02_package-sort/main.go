package main

import (
	"fmt"
	"sort"
)

func main() {
	type people []string
	studyGroup := people{"Zeno", "John", "Al", "Jenny"}
	fmt.Println(studyGroup)
	sort.Strings(studyGroup)
	fmt.Println(studyGroup)
}
