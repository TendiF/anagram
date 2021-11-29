package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	data := make(map[string][]string)
	x := []string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"}

	for _, v := range x {
		data[SortString(v)] = append(data[SortString(v)], v)
	}

	fmt.Println(data)
}

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
