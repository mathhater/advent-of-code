package day09

import (
	"bufio"
	"strconv"

	"github.com/mathhater/advent-of-code/utils"
)

func getBlocks(S string) []int {
	var blocks []int
	id := 0
	slice := []byte(S)
	for i := 0; i < len(slice); i++ {
		cnt := utils.ParseInt(string(slice[i]))
		if i % 2 == 0 {
			for j := 0; j < cnt; j++ {
				blocks = append(blocks, id)
			}
			id++
		} else {
			for j := 0; j < cnt; j++ {
				blocks = append(blocks, -1)
			}
		}
	}
	return blocks
}

func Day09() {
	inputFile := utils.OpenFile("2024/day09/day9.in")
	outputFile := utils.CreateFile("2024/day09/day9.out")
	defer inputFile.Close()
	defer outputFile.Close()
	reader := bufio.NewReader(inputFile)
	writer := bufio.NewWriter(outputFile)

	var S string
	for {
		line, err := utils.ReadLine(reader)
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			panic(err)
		}
		S = line
	}
	N := len(S)

	// part 1
	blocks := getBlocks(S)
	res := 0
	right := len(blocks) - 1
	for i := 0; i < len(blocks); i++ {
		if blocks[i] == -1 {
			for blocks[right] == -1 {
				right--
			}
			if i >= right {
				break
			}
			blocks[i] = blocks[right]
			blocks[right] = -1
		}
	}
	for i := 0; i < len(blocks) && blocks[i] != -1; i++ {
		res += i * blocks[i]
	}
	utils.WriteLine(writer, strconv.Itoa(res))

	// part 2
	S2 := S
	blocks2 := getBlocks(S)
	space := make([]int, N)
	for i := 0; i < N; i++ {
		if i % 2 == 1 {
			space[i] = utils.ParseInt(string(S[i]))
		}
	}
	start := make([]int, N)
	for i, j := 0, 0; i < len(start); i++ {
		if i % 2 == 0 {
			for {
				if blocks2[j] == i / 2 {
					start[i] = j
					break
				}
				j++
			}
		}
	}
	for i := 0; i < len(start) - 1; i++ {
		if i % 2 == 1 {
			start[i] = start[i + 1] - space[i]
		}
	}
	for i := N - 1; i >= 0; i-- {
		if i % 2 == 1 {
			continue
		}
		a := utils.ParseInt(string(S2[i]))
		for j := 0; j < i; j++ {
			if j % 2 == 0 {
				continue
			}
			b := space[j]
			if a <= b {
				for k := 0; k < a; k++ {
					blocks2[start[j] + k] = blocks2[start[i] + k]
					blocks2[start[i] + k] = -1
				}
				start[j] += a
				space[j] -= a
				break
			}
		}
	}
	res2 := 0
	for i := 0; i < len(blocks2); i++ {
		if blocks2[i] != -1 {
			res2 += i * blocks2[i]
		}
	}
	utils.WriteLine(writer, strconv.Itoa(res2))
	writer.Flush()
}