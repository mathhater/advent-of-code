package day03

import (
	"bufio"
	"strconv"

	"github.com/mathhater/advent-of-code/utils"
)

func byteToInt(b byte) int {
	return int(b - '0')
}

func pow10(N int) int {
	res := 1
	for i := 0; i < N; i++ {
		res *= 10
	}
	return res
}

func getJoltage(bank string, size int, pos int) int {
	if size <= 0 || pos + size > len(bank) {
		return 0
	}
	tmp := byteToInt(bank[pos])
	cur := pos
	for i := pos; i <= len(bank) - size; i++ {
		if tmp < byteToInt(bank[i]) {
			tmp = byteToInt(bank[i])
			cur = i
		}
	}
	return tmp * pow10(size - 1) + getJoltage(bank, size - 1, cur + 1)
}

func Day03() {
	inputFile := utils.OpenFile("2025/day03/day03.in")
	outputFile := utils.CreateFile("2025/day03/day03.out")
	defer inputFile.Close()
	defer outputFile.Close()
	reader := bufio.NewReader(inputFile)
	writer := bufio.NewWriter(outputFile)
	
	var banks []string
	for {
		line, err := utils.ReadLine(reader)
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			panic(err)
		}
		banks = append(banks, line)
	}

	// part 1
	res1 := 0
	for i := 0; i < len(banks); i++ {
		res1 += getJoltage(banks[i], 2, 0)
	}
	utils.WriteLine(writer, strconv.Itoa(res1))

	// part 2
	res2 := 0
	for i := 0; i < len(banks); i++ {
		res2 += getJoltage(banks[i], 12, 0)
	}
	utils.WriteLine(writer, strconv.Itoa(res2))
	writer.Flush()
}