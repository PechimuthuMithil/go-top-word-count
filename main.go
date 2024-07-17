package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	words := make(map[string]int)

	scan.Split(bufio.ScanWords)

	for scan.Scan() {
		words[scan.Text()]++
	}

	type key_val struct {
		key   string
		value int
	}

	var ss []key_val

	for key, value := range words {
		ss = append(ss, key_val{
			key:   key,
			value: value,
		})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].value > ss[j].value
	})

	fmt.Println(len(ss), "unique words found")

	podium_len := 10
	// var copy int
	// copy = podium_len
	// cols := 0

	// for copy > 0 {
	// 	copy = copy/10
	// 	cols += 1
	// }

	fmt.Println("---TOP", podium_len, "WORDS---")
	for index, s := range ss[:podium_len] {
		fmt.Printf("[%2d] %4s appears %2d times\n", index+1, s.key, s.value)
		// fmt.Println("[", index+1, "]", s.key, "appears", s.value, "times")
	}
}
