package day12

import (
	"bufio"
	"container/list"
	"strconv"

	"github.com/mathhater/advent-of-code/utils"
)

type Point struct {
	y int
	x int
}

func getPrice1(g [][]byte, visit [][]bool, dir [][]int, y, x, H, W int) int {
	Q := list.New()
	Q.PushBack(Point{y, x})
	visit[y][x] = true
	sides := 0
	area := 0
	for Q.Len() > 0 {
		p := Q.Remove(Q.Front()).(Point)
		area++
		for i := 0; i < len(dir); i++ {
			ty := p.y + dir[i][0]
			tx := p.x + dir[i][1]
			if ty < 0 || ty >= H || tx < 0 || tx >= W {
				sides++
				continue
			}
			if g[ty][tx] == g[y][x] {
				if visit[ty][tx] {
					continue
				}
				Q.PushBack(Point{ty, tx})
				visit[ty][tx] = true
			} else {
				sides++
			}
		}
	}
	return area * sides
}

func getPrice2(g [][]byte, visit [][]bool, dir [][]int, y, x, H, W int) int {
	edge := make([][][]bool, len(dir))
	for i := 0; i < len(edge); i++ {
		edge[i] = make([][]bool, H)
		for j := 0; j < H; j++ {
			edge[i][j] = make([]bool, W)
		}
	}
	Q := []Point{{y, x}}
	visit[y][x] = true
	sides := 0
	area := 1
	for len(Q) > 0 {
		var tmpQ []Point
		for i := 0; i < len(Q); i++ {
			p := Q[i]
			for j := 0; j < len(dir); j++ {
				ty := p.y + dir[j][0]
				tx := p.x + dir[j][1]
				if ty < 0 || ty >= H || tx < 0 || tx >= W || g[ty][tx] != g[y][x] {
					edge[j][p.y][p.x] = true
					sides++
					rdy, rdx := dir[j][1], dir[j][0]
					tty := p.y + rdy
					ttx := p.x + rdx
					
					if 0 <= tty && tty < H && 0 <= ttx && ttx < W && edge[j][tty][ttx] {
						sides--
					}
					tty = p.y + rdy * -1
					ttx = p.x + rdx * -1
					if 0 <= tty && tty < H && 0 <= ttx && ttx < W && edge[j][tty][ttx] {
						sides--
					}
					continue
				}
				if visit[ty][tx] {
					continue
				}
				tmpQ = append(tmpQ, Point{ty, tx})
				visit[ty][tx] = true
				area++
			}
		}
		Q = tmpQ
	}
	return area * sides
}

func Day12() {
	inputFile := utils.OpenFile("2024/day12/day12.in")
	outputFile := utils.CreateFile("2024/day12/day12.out")
	defer inputFile.Close()
	defer outputFile.Close()
	reader := bufio.NewReader(inputFile)
	writer := bufio.NewWriter(outputFile)

	var g [][]byte
	for {
		line, err := utils.ReadLine(reader)
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			panic(err)
		}
		tmp := []byte(line)
		g = append(g, tmp)
	}
	H := len(g)
	W := len(g[0])
	dir := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	// part 1
	res := 0
	visit := make([][]bool, H)
	for i := 0; i < H; i++ {
		visit[i] = make([]bool, W)
	}
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if !visit[i][j] {
				res += getPrice1(g, visit, dir, i, j, H, W)
			}
		}
	}
	utils.WriteLine(writer, strconv.Itoa(res))

	// part 2
	res2 := 0
	visit2 := make([][]bool, H)
	for i := 0; i < H; i++ {
		visit2[i] = make([]bool, W)
	}
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if !visit2[i][j]  {
				res2 += getPrice2(g, visit2, dir, i, j, H, W)
			}
		}
	}
	utils.WriteLine(writer, strconv.Itoa(res2))
	writer.Flush()
}