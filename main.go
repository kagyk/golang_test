package main

import (
	"fmt"
	"sort"
	"strings"
)

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	fmt.Println("Отсортированная строка", s)
	return strings.Join(s, "")
}

func main() {

	var str string = "qqqqqQQQQaaaahhhkkklll"
	var result strings.Builder
	fmt.Println("Входная строка: ", str)
	var count = 0
	str = SortString(str)
	str_len := len(str)
	var c = str[0]
	for i := 0; i < str_len; i++ {

		if c == str[i] {

			count++
			c = str[i]

		} else {

			fmt.Fprintf(&result, "%d%c", count, c)
			c = str[i]
			count = 1

		}

	}
	fmt.Fprintf(&result, "%d%c", count, c)
	fmt.Println("Результат =", result.String())

}
