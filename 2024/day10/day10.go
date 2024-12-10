package day10

import (
	"bufio"
	"container/list"
	"strconv"

	"github.com/mathhater/advent-of-code/utils"
)

type Point struct {
	x int
	y int
}

func getScore(g, dir [][]int, y, x, H, W int) int {
	res := 0
	list := list.New()
	list.PushBack(Point{x, y})
	visit := make([][]bool, H)
	for i := 0; i < len(visit); i++ {
		visit[i] = make([]bool, W)
	}
	visit[y][x] = true
	for list.Len() > 0 {
		p := list.Remove(list.Front()).(Point)
		for i := 0; i < len(dir); i++ {
			ty := p.y + dir[i][0]
			tx := p.x + dir[i][1]
			if ty < 0 || ty >= H || tx < 0 || tx >= W || visit[ty][tx] || g[ty][tx] != g[p.y][p.x] + 1 {
				continue
			}
			if g[ty][tx] == 9 {
				visit[ty][tx] = true
				res++
				continue
			}
			list.PushBack(Point{tx, ty})
			visit[ty][tx] = true
		}
	}
	return res
}

func getScore2(g, dir [][]int, y, x, H, W int) int {
	res := 0
	list := list.New()
	list.PushBack(Point{x, y})
	visit := make([][]bool, H)
	for i := 0; i < len(visit); i++ {
		visit[i] = make([]bool, W)
	}
	visit[y][x] = true
	dp := make([][]int, H)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, W)
	}
	dp[y][x] = 1
	for list.Len() > 0 {
		p := list.Remove(list.Front()).(Point)
		for i := 0; i < len(dir); i++ {
			ty := p.y + dir[i][0]
			tx := p.x + dir[i][1]
			if ty < 0 || ty >= H || tx < 0 || tx >= W || g[ty][tx] != g[p.y][p.x] + 1 {
				continue
			}
			dp[ty][tx] += dp[p.y][p.x]
			if g[ty][tx] == 9 {
				continue
			}
			if !visit[ty][tx] {
				list.PushBack(Point{tx, ty})
			}
			visit[ty][tx] = true
		}
	}
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if g[i][j] == 9 {
				res += dp[i][j]
			}
		}
	}
	return res
}

func Day10() {
	inputFile := utils.OpenFile("2024/day10/day10.in")
	outputFile := utils.CreateFile("2024/day10/day10.out")
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
	H := len(S)
	W := len(S[0])
	var g [][]int
	for i := 0; i < len(S); i++ {
		var tmp []int
		for j := 0; j < len(S[i]); j++ {
			tmp = append(tmp, utils.ParseInt(string(S[i][j])))
		}
		g = append(g, tmp)
	}

	// part 1
	res := 0
	dir := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if g[i][j] == 0 {
				res += getScore(g, dir, i, j, H, W)
			}
		}
	}
	utils.WriteLine(writer, strconv.Itoa(res))

	// part 2
	res2 := 0
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if g[i][j] == 0 {
				res2 += getScore2(g, dir, i, j, H, W)
			}
		}
	}
	utils.WriteLine(writer, strconv.Itoa(res2))
	writer.Flush()
}