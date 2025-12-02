package day02

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/mathhater/advent-of-code/utils"
)

func getWrongId1(id string) int {
	if len(id) % 2 != 0 {
		return 0
	}
	for i := 0; i < len(id) / 2; i++ {
		if id[i] != id[len(id) / 2 + i] {
			return 0
		}
	}
	return utils.ParseInt(id)
}

func getWrongId2(id string) int {
	for i := 1; i < len(id); i++ {
		if len(id) % i != 0 {
			continue
		}

		flag := true
		for j := 0; j < len(id) / i; j++ {
			if id[:i] != id[j * i : (j + 1) * i] {
				flag = false
				break
			}
		}

		if flag {
			return utils.ParseInt(id)
		}
	}
	return 0
}

func Day02() {
	inputFile := utils.OpenFile("2025/day02/day02.in")
	outputFile := utils.CreateFile("2025/day02/day02.out")
	defer inputFile.Close()
	defer outputFile.Close()
	reader := bufio.NewReader(inputFile)
	writer := bufio.NewWriter(outputFile)

	line, err := utils.ReadLine(reader)
	if err != nil {
		panic(err)
	}
	S := strings.Split(line, ",")

	// part 1
	res1 := 0
	for i := 0; i < len(S); i++ {
		tmpS := strings.Split(S[i], "-")
		for j := utils.ParseInt(tmpS[0]); j <= utils.ParseInt(tmpS[1]); j++ {
			res1 += getWrongId1(strconv.Itoa(j))
		}
	}
	utils.WriteLine(writer, strconv.Itoa(res1))

	// part 2
	res2 := 0
	for i := 0; i < len(S); i++ {
		tmpS := strings.Split(S[i], "-")
		for j := utils.ParseInt(tmpS[0]); j <= utils.ParseInt(tmpS[1]); j++ {
			res2 += getWrongId2(strconv.Itoa(j))
		}
	}
	utils.WriteLine(writer, strconv.Itoa(res2))
	writer.Flush()
}