package day01

import (
	"bufio"
	"strconv"

	"github.com/mathhater/advent-of-code/utils"
)

func Day01() {
	inputFile := utils.OpenFile("2025/day01/day01.in")
	outputFile := utils.CreateFile("2025/day01/day01.out")
	defer inputFile.Close()
	defer outputFile.Close()
	reader := bufio.NewReader(inputFile)
	writer := bufio.NewWriter(outputFile)

	var dial int
	var S []string
	var BIG_INT = 1000000000
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
	var res1 = 0
	dial = 50
	for i := 0; i < len(S); i++ {
		var rotation = utils.ParseInt(S[i][1:])
		if S[i][0] == 'L' {
			for j := 0; j < rotation; j++ {
				dial--
			}
		} else {
			for j := 0; j < rotation; j++ {
				dial++
			}
		}
		if (dial + BIG_INT) % 100 == 0 {
			res1++
		}
	}
	utils.WriteLine(writer, strconv.Itoa(res1))

	// part 2
	var res2 = 0
	dial = 50
	for i := 0; i < len(S); i++ {
		var rotation = utils.ParseInt(S[i][1:])
		if S[i][0] == 'L' {
			for j := 0; j < rotation; j++ {
				dial--
				if (dial + BIG_INT) % 100 == 0 {
					res2++
				}
			}
		} else {
			for j := 0; j < rotation; j++ {
				dial++
				if (dial + BIG_INT) % 100 == 0 {
					res2++
				}
			}
		}
	}

	utils.WriteLine(writer, strconv.Itoa(res2))
	writer.Flush()
}