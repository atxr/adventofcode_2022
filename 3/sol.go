package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func contains(s string, a rune) bool {
	for _, e := range(s) {
		if e == a {
			return true
		}
	}
	return false
}

func getCommon2(a, b string) []rune {
	res := make([]rune, 0)
	for _, l_a := range(a) {
		if contains(b, l_a) && !contains(string(res), l_a) {
			res = append(res, l_a)
		}
	}
	return res
}

func getCommon3(a, b, c string) []rune {
	return getCommon2(string(getCommon2(a, b)), c)
}

func split(l string) (string, string) {
	n := int(len(l)/2)
	return l[:n], l[n:]
}

func getPriority(r rune) int {
	if unicode.IsLower(r) {
		return int(r - 'a') + 1
	}
	return int(r - 'A') + 27
}

func main() {
	src, err := os.Open("input")
	defer src.Close()
	check(err)
	scanner := bufio.NewScanner(src)

	score1, score2, k := 0, 0, 0
	buf := make([]string, 3)
	for scanner.Scan() {
		buf[k] = scanner.Text()
		a, b := split(buf[k])
		k = (k+1) % 3
		err := getCommon2(a, b)
		if len(err) != 1 {
			log.Fatal("score 1 ", err)
		}
		score1 += getPriority(err[0])

		if (k % 3 != 0) {
			continue
		}

		badge := getCommon3(buf[0], buf[1], buf[2])
		if len(badge) != 1 {
			log.Fatal("score 2 ", buf)
		}
		score2 += getPriority(badge[0])
	}

	fmt.Println(score1)
	fmt.Println(score2)
}
