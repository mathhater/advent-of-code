package day07

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/mathhater/advent-of-code/utils"
)

func setResults(res *[]int, b []int, cur int, idx int, idxres *int) {
	if idx == len(b) {
		(*res)[*idxres] = cur
		*idxres++
		return
	}
	setResults(res, b, cur+b[idx], idx+1, idxres)
	setResults(res, b, cur*b[idx], idx+1, idxres)
}

func setResults2(res *[]int, b []int, cur int, idx int, idxres *int) {
	if idx == len(b) {
		(*res)[*idxres] = cur
		*idxres++
		return
	}
	setResults2(res, b, cur+b[idx], idx+1, idxres)
	setResults2(res, b, cur*b[idx], idx+1, idxres)
	digit := 1
	x := b[idx]
	for {
		digit *= 10
		x /= 10
		if x == 0 {
			break
		}
	}
	setResults2(res, b, cur*digit+b[idx], idx+1, idxres)
}

func equationCheck(a int, b []int) bool {
	var idxres = 0
	res := make([]int, (1 << len(b)))
	setResults(&res, b, 0, 0, &idxres)
	for i := 0; i < len(res); i++ {
		if res[i] == a {
			return true
		}
	}
	return false
}

func equationCheck2(a int, b []int) bool {
	var idxres = 0
	resSize := 1
	for i := 0; i < len(b); i++ {
		resSize *= 3
	}
	res := make([]int, resSize)
	setResults2(&res, b, b[0], 1, &idxres)
	for i := 0; i < len(res); i++ {
		if res[i] == a {
			return true
		}
	}
	return false
}

func Day07() {
	inputFile := utils.OpenFile("2024/day07/day7.in")
	outputFile := utils.CreateFile("2024/day07/day7.out")
	defer inputFile.Close()
	defer outputFile.Close()
	reader := bufio.NewReader(inputFile)
	writer := bufio.NewWriter(outputFile)

	var A []int
	var B [][]int
	idxB := 0
	for {
		line, err := utils.ReadLine(reader)
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			panic(err)
		}

		splittedLine := strings.Split(line, " ")
		A = append(A, utils.ParseInt(strings.TrimRight(splittedLine[0], ":")))
		var tmpB []int
		for i := 1; i < len(splittedLine); i++ {
			tmpB = append(tmpB, utils.ParseInt(splittedLine[i]))
		}
		B = append(B, tmpB)
		idxB++
	}

	// part 1
	res := 0
	for i := 0; i < len(A); i++ {
		if equationCheck(A[i], B[i]) {
			res += A[i]
		}
	}
	utils.WriteLine(writer, strconv.Itoa(res))

	// part 2
	res2 := 0
	for i := 0; i < len(A); i++ {
		if equationCheck2(A[i], B[i]) {
			res2 += A[i]
		}
	}
	utils.WriteLine(writer, strconv.Itoa(res2))
	writer.Flush()
}
