package day11

import (
	"bufio"
	"math/big"
	"strconv"
	"strings"

	"github.com/mathhater/advent-of-code/utils"
)

func getBlink(dp []map[string]int, num big.Int, times int) int {
	if times == 0 {
		return 1
	}
	key := num.Text(10)
	if _, exists := dp[times][key]; exists {
		return dp[times][key]
	}

	zero := big.NewInt(0)
	multiplier := big.NewInt(2024)
	if num.Cmp(zero) == 0 {
		var bigint big.Int
		bigint.SetString("1", 10)
		dp[times][key] = getBlink(dp, bigint, times - 1)
		return dp[times][key]
	}

	digits := num.Text(10)
	length := len(digits)
	if length % 2 == 0 {
		mid := length / 2
		var bigint1, bigint2 big.Int
		bigint1.SetString(digits[0:mid], 10)
		bigint2.SetString(digits[mid:], 10)
		dp[times][key] = getBlink(dp, bigint1, times - 1) + getBlink(dp, bigint2, times - 1)
		return dp[times][key]
	}
	var bigint big.Int
	bigint.Mul(&num, multiplier)
	dp[times][key] = getBlink(dp, bigint, times - 1)
	return dp[times][key]
}

func Day11() {
	inputFile := utils.OpenFile("2024/day11/day11.in")
	outputFile := utils.CreateFile("2024/day11/day11.out")
	defer inputFile.Close()
	defer outputFile.Close()
	reader := bufio.NewReader(inputFile)
	writer := bufio.NewWriter(outputFile)

	var S[] string
	for {
		line, err := utils.ReadLine(reader)
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			panic(err)
		}
		S = strings.Split(line, " ")
	}
	A := make([]big.Int, len(S))
	for i := 0; i < len(S); i++ {
		A[i].SetString(S[i], 10)
	}
	dp := make([]map[string]int, 76)
	for i := 0; i < len(dp); i++ {
		dp[i] = make(map[string]int)
	}

	// part 1
	res := 0
	for i := 0; i < len(A); i++ {
		res += getBlink(dp, A[i], 25)
	}
	utils.WriteLine(writer, strconv.Itoa(res))

	// part 2
	res2 := 0
	for i := 0; i < len(A); i++ {
		res2 += getBlink(dp, A[i], 75)
	}
	utils.WriteLine(writer, strconv.Itoa(res2))
	writer.Flush()
}