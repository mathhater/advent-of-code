package day02

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/mathhater/advent-of-code/utils"
)

func isSafe(A []int) bool {
	if A[0] < A[1] {
		cnt := 0
		for i := 1; i < len(A); i++ {
			if A[i - 1] > A[i] {
				break
			}
			if A[i] - A[i - 1] > 3 {
				break
			}
			if A[i] == A[i - 1] {
				break
			}
			cnt++
		}
		return cnt == len(A) - 1
	} else if A[0] > A[1] {
		cnt := 0
		for i := 1; i < len(A); i++ {
			if A[i - 1] < A[i] {
				break
			}
			if A[i - 1] - A[i] > 3 {
				break
			}
			if A[i] == A[i - 1] {
				break
			}
			cnt++
		}
		return cnt == len(A) - 1
	}
	return false
}

func Day02() {
	inputFile := utils.OpenFile("2024/day02/day2.in")
	outputFile := utils.CreateFile("2024/day02/day2.out")
	defer inputFile.Close()
	defer outputFile.Close()
	reader := bufio.NewReader(inputFile)
	writer := bufio.NewWriter(outputFile)

	var A [][]int
	for {
		line, err := utils.ReadLine(reader)
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			panic(err)
		}

		splittedLine := strings.Split(line, " ")
		var tmpA []int
		for i := 0; i < len(splittedLine); i++ {
			tmpA = append(tmpA, utils.ParseInt(splittedLine[i]))
		}
		A = append(A, tmpA)
	}

	// part 1
	res := 0
	for i := 0; i < len(A); i++ {
		if isSafe(A[i]) {
			res++
		}
	}
	utils.WriteLine(writer, strconv.Itoa(res))

	// part 2
	res2 := 0
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[i]); j++ {
			var tmpA []int
			tmpA = append(tmpA, A[i][0 : j]...)
			tmpA = append(tmpA, A[i][j + 1 :]...)
			if isSafe(tmpA) {
				res2++
				break
			}
		}
	}
	utils.WriteLine(writer, strconv.Itoa(res2))
	writer.Flush()
}