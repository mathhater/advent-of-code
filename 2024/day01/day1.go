package day01

import (
	"bufio"
	"sort"
	"strconv"
	"strings"

	"github.com/mathhater/advent-of-code/utils"
)

func Day01() {
	inputFile := utils.OpenFile("2024/day01/day1.in")
	outputFile := utils.CreateFile("2024/day01/day1.out")
	defer inputFile.Close()
	defer outputFile.Close()
	reader := bufio.NewReader(inputFile)
	writer := bufio.NewWriter(outputFile)

	var A, B []int
	for {
		line, err := utils.ReadLine(reader)
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			panic(err)
		}

		splittedLine := strings.Split(line, "   ")
		a := utils.ParseInt(splittedLine[0])
		b := utils.ParseInt(splittedLine[1])
		A = append(A, int(a))
		B = append(B, int(b))
	}
	sort.Ints(A)
	sort.Ints(B)

	// part 1
	res := 0
	for i := 0; i < len(A); i++ {
		res += utils.Abs(B[i] - A[i])
	}
	utils.WriteLine(writer, strconv.Itoa(res))

	// part 2
	res2 := 0
	var cnt = make(map[int]int)
	for i := 0; i < len(A); i++ {
		cnt[B[i]]++
	}
	for i := 0; i < len(A); i++ {
		res2 += A[i] * cnt[A[i]]
	}
	utils.WriteLine(writer, strconv.Itoa(res2))
	writer.Flush()
}