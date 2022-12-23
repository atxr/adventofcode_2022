package main 

import (
	"fmt"
	"bufio"
	"os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func getScore1(op, play byte) uint8 {
	u_op, u_play := uint8(op - 'A'), uint8(play - 'X')
	return u_play+1 + 3 * ((4 + u_play - u_op) % 3)
}

func getScore2(op, res byte) uint8 {
	play := ((op - 'A') + (res - 'X' + 2)) % 3 + 'X'
	return getScore1(op, play)
}

func main() {
	src, err := os.Open("input")
	defer src.Close()
	check(err)
	scanner := bufio.NewScanner(src)

	score1 := uint64(0)
	score2 := uint64(0)
	for scanner.Scan() {
		line := scanner.Text()
		score1 += uint64(getScore1(line[0], line[2]))
		score2 += uint64(getScore2(line[0], line[2]))
	}

	fmt.Println(score1)
	fmt.Println(score2)
}
