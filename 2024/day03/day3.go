package day03

import (
	"bufio"
	"strconv"

	"github.com/mathhater/advent-of-code/utils"
)

func Day03() {
	inputFile := utils.OpenFile("2024/day03/day3.in")
	outputFile := utils.CreateFile("2024/day03/day3.out")
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

	res := 0
	regexNumber := utils.GetRegex(`[0-9]+`)
	regexMull := utils.GetRegex(`mul\((\d+),\s*(\d+)\)`)
	for i := 0; i < len(S); i++ {
		parsedString := regexMull.FindAllString(S[i], -1)
		for j := 0; j < len(parsedString); j++ {
			result := regexNumber.FindAllString(parsedString[j], -1)
			a := utils.ParseInt(result[0])
			b := utils.ParseInt(result[1])
			res += a * b
		}
	}
	utils.WriteLine(writer, strconv.Itoa(res))

	res2 := 0
	flag := true
	regexMull2 := utils.GetRegex(`mul\((\d+),\s*(\d+)\)|do\(\)|don't\(\)`)
	for i := 0; i < len(S); i++ {
		parsedString := regexMull2.FindAllString(S[i], -1)
		for j := 0; j < len(parsedString); j++ {
			if parsedString[j] == "do()" {
				flag = true
				continue
			}
			if parsedString[j] == "don't()" {
				flag = false
				continue
			}
			if !flag {
				continue
			}
			result := regexNumber.FindAllString(parsedString[j], -1)
			a := utils.ParseInt(result[0])
			b := utils.ParseInt(result[1])
			res2 += a * b
		}
	}
	utils.WriteLine(writer, strconv.Itoa(res2))
	writer.Flush()
}