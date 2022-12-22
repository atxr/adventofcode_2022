package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func max(l []uint64) (uint64, int) {
	var m uint64
	var im int
	for i, cal := range l {
	    if i==0 || cal > m {
		m = cal
		im = i
	    }
	}
	return m, im
}

func sort(l []uint64) []uint64 {
	sorted := l
	for k := range sorted {
		m, im := max(sorted[k:])
		sorted[im + k] = sorted[k]
		sorted[k] = m
	}
	return sorted
}

func sum(l []uint64) uint64 {
	s := uint64(0)
	for _, el := range l {
		s += el
	}
	return s
}

func main() {
	src, err := os.Open("input")
	defer src.Close()
	check(err)
	scanner := bufio.NewScanner(src)

	supplies := make([]uint64, 0)
	s := uint64(0)
	for scanner.Scan() {
		l := scanner.Text()
		if l == "" {
			supplies = append(supplies, s)
			s = 0
		} else {
			i, err := strconv.Atoi(l)
			check(err)
			s += uint64(i)
		}
	}
	supplies = append(supplies, s)
	supplies = sort(supplies)

	fmt.Printf("The strongest elf carries %d calories!\n", supplies[0])
	fmt.Printf("The top 3 elves carry %d calories!\n", sum(supplies[:3]))
}
