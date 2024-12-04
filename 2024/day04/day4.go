package day04

import (
	"bufio"
	"strconv"

	"github.com/mathhater/advent-of-code/utils"
)

func getXMas(S []string, y int, x int, h int, w int) int {
	dir := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1}}
	XMAS := "XMAS"
	res := 0
	for k := 0; k < len(dir); k++ {
		ty := y
		tx := x
		cnt := 0
		for {
			if ty < 0 || ty >= h || tx < 0 || tx >= w {
				break
			}
			if S[ty][tx] != XMAS[cnt] {
				break
			}
			cnt++
			if cnt == len(XMAS) {
				res++
				break
			}
			ty += dir[k][0]
			tx += dir[k][1]
		}
	}
	return res
}

func isMas(S []string, y int, x int) bool {
	if S[y][x] != 'A' {
		return false
	}
	if S[y + 1][x + 1] == 'M' && S[y - 1][x - 1] == 'S' {
		if S[y + 1][x - 1] == 'M' && S[y - 1][x + 1] == 'S' {
			return true
		} else if S[y + 1][x - 1] == 'S' && S[y - 1][x + 1] == 'M' {
			return true
		}
	}
	if S[y + 1][x + 1] == 'S' && S[y - 1][x - 1] == 'M' {
		if S[y + 1][x - 1] == 'M' && S[y - 1][x + 1] == 'S' {
			return true
		} else if S[y + 1][x - 1] == 'S' && S[y - 1][x + 1] == 'M' {
			return true
		}
	}
	return false
}

func Day04() {
	inputFile := utils.OpenFile("2024/day04/day4.in")
	outputFile := utils.CreateFile("2024/day04/day4.out")
	defer inputFile.Close()
	defer outputFile.Close()
	reader := bufio.NewReader(inputFile)
	writer := bufio.NewWriter(outputFile)

	var S []string
	for {
		line, err := utils.ReadLine(reader)
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			panic(err)
		}
		S = append(S, line)
	}

	// part 1
	res := 0
	for i := 0; i < len(S); i++ {
		for j := 0; j < len(S[i]); j++ {
			res += getXMas(S, i, j, len(S), len(S[i]));
		}
	}
	utils.WriteLine(writer, strconv.Itoa(res))

	// part 2
	res2 := 0
	for i := 1; i < len(S) - 1; i++ {
		for j := 1; j < len(S[i]) - 1; j++ {
			if isMas(S, i, j) {
				res2++
			}
		}
	}
	utils.WriteLine(writer, strconv.Itoa(res2))
	writer.Flush()
}