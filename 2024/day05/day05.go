package day05

import (
	"bufio"
	"container/list"
	"strconv"
	"strings"

	"github.com/mathhater/advent-of-code/utils"
)

func addRule(A map[int]map[int]bool, a int, b int) {
	if _, exists := A[a]; !exists {
		A[a] = make(map[int]bool)
	}
	A[a][b] = true
}

func Day05() {
	inputFile := utils.OpenFile("2024/day05/day5.in")
	outputFile := utils.CreateFile("2024/day05/day5.out")
	defer inputFile.Close()
	defer outputFile.Close()
	reader := bufio.NewReader(inputFile)
	writer := bufio.NewWriter(outputFile)

	A := make(map[int]map[int]bool)
	for {
		line, err := utils.ReadLine(reader)
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			panic(err)
		}
		if line == "" {
			break
		}

		splittedLine := strings.Split(line, "|")
		a := utils.ParseInt(splittedLine[0])
		b := utils.ParseInt(splittedLine[1])
		addRule(A, a, b)
	}
	
	var B [][]int
	for {
		line, err := utils.ReadLine(reader)
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			panic(err)
		}

		splittedLine := strings.Split(line, ",")
		var tmpB []int
		for i := 0; i < len(splittedLine); i++ {
			x := utils.ParseInt(splittedLine[i])
			tmpB = append(tmpB, x)
		}
		B = append(B, tmpB)
	}

	// part 1
	res := 0
	flag := make([]bool, len(B))
	for i := 0; i < len(B); i++ {
		flag[i] = true
		for j := 0; j < len(B[i]); j++ {
			for k := j + 1; k < len(B[i]); k++ {
				if !A[B[i][j]][B[i][k]] {
					flag[i] = false
				}
			}
		}
		if flag[i] {
			res += B[i][len(B[i]) / 2];
		}
	}
	utils.WriteLine(writer, strconv.Itoa(res))

	// part 2
	res2 := 0
	for i := 0; i < len(B); i++ {
		if !flag[i] {
			g := make(map[int][]int)
			cnt := make(map[int]int)
			for j := 0; j < len(B[i]); j++ {
				for k := 0; k < len(B[i]); k++ {
					if j == k {
						continue
					}
					if A[B[i][j]][B[i][k]] {
						g[B[i][k]] = append(g[B[i][k]], B[i][j])
						cnt[B[i][j]]++
					} else {
						if _, exists := cnt[B[i][j]]; !exists {
							cnt[B[i][j]] = 0
						}
					}
				}
			}
			
			var tmpB []int
			list := list.New()
			for key, value := range cnt {
				if value == 0 {
					list.PushBack(key)
				}
			}
			for {
				if list.Len() == 0 {
					break
				}
				node := list.Front()
				x := node.Value.(int)
				list.Remove(node)
				tmpB = append(tmpB, x)
				for i := 0; i < len(g[x]); i++ {
					cnt[g[x][i]]--
					if cnt[g[x][i]] == 0 {
						list.PushBack(g[x][i])
					}
				}
			}
			res2 += tmpB[len(tmpB) / 2]
		}
	}
	utils.WriteLine(writer, strconv.Itoa(res2))
	writer.Flush()
}