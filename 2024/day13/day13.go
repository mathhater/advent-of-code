package day13

import (
	"bufio"
	"math"
	"strconv"
	"strings"

	"github.com/mathhater/advent-of-code/utils"
)

func parseButton(s string) (int, int) {
	//Button A: X+94, Y+34
	splittedStrings := strings.Split(s, " ")
	length := len(splittedStrings[2])
	a := utils.ParseInt(splittedStrings[2][2:(length-1)])
	b := utils.ParseInt(splittedStrings[3][2:])
	return a, b
}

func parsePrize(s string) (int, int) {
	//Prize: X=8400, Y=5400
	splittedStrings := strings.Split(s, " ")
	length := len(splittedStrings[1])
	a := utils.ParseInt(splittedStrings[1][2:(length-1)])
	b := utils.ParseInt(splittedStrings[2][2:])
	return a, b
}

func getSolution(a, b, c, d, p, q int) (int, int) {
	D := a * d - b * c
	if D == 0 {
		return 0, 0
	}
	if (d * p  - b * q) % D != 0 || (a * q - c * p) % D != 0 {
		return 0, 0
	}
	x := (d * p  - b * q) / D
	y := (a * q - c * p) / D
	return x, y
}

func Day13() {
	inputFile := utils.OpenFile("2024/day13/day13.in")
	outputFile := utils.CreateFile("2024/day13/day13.out")
	defer inputFile.Close()
	defer outputFile.Close()
	reader := bufio.NewReader(inputFile)
	writer := bufio.NewWriter(outputFile)

	var matrix [][][]int
	for {
		line1, err := utils.ReadLine(reader)
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			panic(err)
		}
		line2, err := utils.ReadLine(reader)
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			panic(err)
		}
		line3, err := utils.ReadLine(reader)
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			panic(err)
		}
		_, _ = utils.ReadLine(reader)
		tmpMatrix := make([][]int, 2)
		tmpMatrix[0] = make([]int, 3)
		tmpMatrix[1] = make([]int, 3)
		tmpMatrix[0][0], tmpMatrix[1][0] = parseButton(line1)
		tmpMatrix[0][1], tmpMatrix[1][1] = parseButton(line2)
		tmpMatrix[0][2], tmpMatrix[1][2] = parsePrize(line3)
		matrix = append(matrix, tmpMatrix)
	}

	// part 1
	res := 0
	for i := 0; i < len(matrix); i++ {
		a := matrix[i][0][0]
		b := matrix[i][0][1]
		c := matrix[i][0][2]
		d := matrix[i][1][0]
		e := matrix[i][1][1]
		f := matrix[i][1][2]
		tmp := math.MaxInt32
		for x := 0; a * x <= c && d * x <= f; x++ {
			if (c - a * x) % b == 0 && (f - d * x) % e == 0 {
				y := (c - a * x) / b
				if a * x + b * y == c && d * x + e * y == f && tmp > 3 * x + y {
					tmp = 3 * x + y
				}
			}
		}
		if tmp < math.MaxInt32 {
			res += tmp
		}
	}
	utils.WriteLine(writer, strconv.Itoa(res))

	// part 2
	res2 := 0
	added := 10000000000000
	for i := 0; i < len(matrix); i++ {
		a := matrix[i][0][0]
		b := matrix[i][0][1]
		p := matrix[i][0][2] + added
		c := matrix[i][1][0]
		d := matrix[i][1][1]
		q := matrix[i][1][2] + added
		x, y := getSolution(a, b, c, d, p, q)
		if x < 0 || y < 0 {
			continue
		}
		res2 += 3 * x + y
	}
	utils.WriteLine(writer, strconv.Itoa(res2))
	writer.Flush()
}