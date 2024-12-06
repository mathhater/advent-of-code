package day06

import (
	"bufio"
	"strconv"

	"github.com/mathhater/advent-of-code/utils"
)

func getNext(g []string, dir [][]int, y int, x int, d int, H int, W int) (int, int, int) {
	ty := y + dir[d][0]
	tx := x + dir[d][1]
	if ty < 0 || ty >= H || tx < 0 || tx >= W {
		return ty, tx, d
	}
	if g[ty][tx] == '#' {
		d = (d + 1) % 4
		return y, x, d
	}
	return ty, tx, d
}

func loopCheck(g []string, dir [][]int, y int, x int, d int, H int, W int) bool {
	visit := make([][][]bool, H)
	for i := 0; i < len(visit); i++ {
		visit[i] = make([][]bool, W)
		for j := 0; j < len(visit[i]); j++ {
			visit[i][j] = make([]bool, 4)
		}
	}
	visit[y][x][d] = true
	for {
		ty, tx, td := getNext(g, dir, y, x, d, H, W)
		if ty < 0 || ty >= H || tx < 0 || tx >= W {
			break
		}
		if visit[ty][tx][td] {
			return true
		}
		visit[ty][tx][td] = true
		y = ty
		x = tx
		d = td
	}
	return false
}

func Day06() {
	inputFile := utils.OpenFile("2024/day06/day6.in")
	outputFile := utils.CreateFile("2024/day06/day6.out")
	defer inputFile.Close()
	defer outputFile.Close()
	reader := bufio.NewReader(inputFile)
	writer := bufio.NewWriter(outputFile)

	var g []string
	for {
		line, err := utils.ReadLine(reader)
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			panic(err)
		}
		g = append(g, line)
	}
	H := len(g)
	W := len(g[0])

	// part 1
	res := 0
	guardSymbol := "^>v<"
	dir := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	var sy, sx, sd int
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			for k := 0; k < len(guardSymbol); k++ {
				if g[i][j] == guardSymbol[k] {
					sy = i
					sx = j
					sd = k
				}
			}
		}
	}
	visit := make([][]bool, H)
	for i := 0; i < len(visit); i++ {
		visit[i] = make([]bool, W)
	}
	y := sy
	x := sx
	d := sd
	visit[y][x] = true
	res++
	for {
		ty, tx, td := getNext(g, dir, y, x, d, H, W)
		if ty < 0 || ty >= H || tx < 0 || tx >= W {
			break
		}
		if !visit[ty][tx] {
			visit[ty][tx] = true
			res++
		}
		y = ty
		x = tx
		d = td
	}
	utils.WriteLine(writer, strconv.Itoa(res))

	// part 2
	res2 := 0
	y = sy
	x = sx
	d = sd
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if g[i][j] == '.' {
				slice := []byte(g[i])
				slice[j] = byte('#')
				g[i] = string(slice)
				if loopCheck(g, dir, y, x, d, H, W) {
					res2++
				}
				slice = []byte(g[i])
				slice[j] = byte('.')
				g[i] = string(slice)
			}
		}
	}
	utils.WriteLine(writer, strconv.Itoa(res2))
	writer.Flush()
}