package day15

import (
	"bufio"
	"strconv"

	"github.com/mathhater/advent-of-code/utils"
)

type Point struct {
	x, y int
}

func moveBox(g [][]byte, dir map[string]Point, cur Point, move byte, H, W int) bool {
	ty := cur.y + dir[string(move)].y
	tx := cur.x + dir[string(move)].x
	if ty < 0 || ty >= H || tx < 0 || tx >= W || g[ty][tx] == '#' {
		return false
	}
	if g[ty][tx] == '.' {
		g[ty][tx] = 'O'
		g[cur.y][cur.x] = '.'
		return true
	}
	if moveBox(g, dir, Point{tx, ty}, move, H, W) {
		g[ty][tx] = 'O'
		g[cur.y][cur.x] = '.'
		return true
	}
	return false
}

func moveRobot(g [][]byte, dir map[string]Point, cur Point, move byte, H, W int) Point {
	ty := cur.y + dir[string(move)].y
	tx := cur.x + dir[string(move)].x
	if ty < 0 || ty >= H || tx < 0 || tx >= W || g[ty][tx] == '#' {
		return cur
	}
	if g[ty][tx] == '.' {
		g[ty][tx] = '@'
		g[cur.y][cur.x] = '.'
		return Point{tx, ty}
	}
	box := Point{tx, ty}
	if moveBox(g, dir, box, move, H, W) {
		g[ty][tx] = '@'
		g[cur.y][cur.x] = '.'
		return Point{tx, ty}
	}
	return cur
}

func moveBox2UD(g [][]byte, dir map[string]Point, cur Point, move byte, H, W int) bool {
	var cx1, cx2, cy1, cy2 int
	if g[cur.y][cur.x] == '[' {
		cx1 = cur.x
		cx2 = cur.x + 1
		cy1 = cur.y
		cy2 = cur.y
	} else {
		cx1 = cur.x - 1
		cx2 = cur.x
		cy1 = cur.y
		cy2 = cur.y
	}
	ty1 := cy1 + dir[string(move)].y
	tx1 := cx1 + dir[string(move)].x
	ty2 := cy2 + dir[string(move)].y
	tx2 := cx2 + dir[string(move)].x
	if ty1 < 0 || ty1 >= H || tx1 < 0 || tx1 >= W || ty2 < 0 || ty2 >= H || tx2 < 0 || tx2 >= W || g[ty1][tx1] == '#' || g[ty2][tx2] == '#' {
		return false
	}
	if g[ty1][tx1] == '.' && g[ty2][tx2] == '.' {
		g[ty1][tx1] = '['
		g[ty2][tx2] = ']'
		g[cy1][cx1] = '.'
		g[cy2][cx2] = '.'
		return true
	}
	if g[ty1][tx1] == '[' && g[ty2][tx2] == ']' {
		if moveBox2UD(g, dir, Point{tx1, ty1}, move, H, W) {
			g[ty1][tx1] = '['
			g[ty2][tx2] = ']'
			g[cy1][cx1] = '.'
			g[cy2][cx2] = '.'
			return true
		}
		return false
	}
	if g[ty1][tx1] == ']' && g[ty2][tx2] == '[' {
		graph := make([][]byte, H)
		for i := 0; i < H; i++ {
			graph[i] = make([]byte, W)
			for j := 0; j < W; j++ {
				graph[i][j] = g[i][j]
			}
		}
		if moveBox2UD(g, dir, Point{tx1, ty1}, move, H, W) && moveBox2UD(g, dir, Point{tx2, ty2}, move, H, W) {
			g[ty1][tx1] = '['
			g[ty2][tx2] = ']'
			g[cy1][cx1] = '.'
			g[cy2][cx2] = '.'
			return true
		}
		for i := 0; i < H; i++ {
			for j := 0; j < W; j++ {
				g[i][j] = graph[i][j]
			}
		}
		return false
	}
	if g[ty1][tx1] == ']' {
		if moveBox2UD(g, dir, Point{tx1, ty1}, move, H, W) {
			g[ty1][tx1] = '['
			g[ty2][tx2] = ']'
			g[cy1][cx1] = '.'
			g[cy2][cx2] = '.'
			return true
		}
		return false
	}
	if g[ty2][tx2] == '[' {
		if moveBox2UD(g, dir, Point{tx2, ty2}, move, H, W) {
			g[ty1][tx1] = '['
			g[ty2][tx2] = ']'
			g[cy1][cx1] = '.'
			g[cy2][cx2] = '.'
			return true
		}
		return false
	}
	return false
}

func moveBox2RL(g [][]byte, dir map[string]Point, cur Point, move byte, H, W int) bool {
	ty := cur.y + dir[string(move)].y
	tx := cur.x + dir[string(move)].x
	if ty < 0 || ty >= H || tx < 0 || tx >= W || g[ty][tx] == '#' {
		return false
	}
	if g[ty][tx] == '.' {
		g[ty][tx] = g[cur.y][cur.x]
		g[cur.y][cur.x] = '.'
		return true
	}
	if moveBox2RL(g, dir, Point{tx, ty}, move, H, W) {
		g[ty][tx] = g[cur.y][cur.x]
		g[cur.y][cur.x] = '.'
		return true
	}
	return false
}

func moveRobot2(g [][]byte, dir map[string]Point, cur Point, move byte, H, W int) Point {
	ty := cur.y + dir[string(move)].y
	tx := cur.x + dir[string(move)].x
	if ty < 0 || ty >= H || tx < 0 || tx >= W || g[ty][tx] == '#' {
		return cur
	}
	if g[ty][tx] == '.' {
		g[cur.y][cur.x] = '.'
		g[ty][tx] = '@'
		return Point{tx, ty}
	}
	if move == '^' || move == 'v' {
		if moveBox2UD(g, dir, Point{tx, ty}, move, H, W) {
			g[cur.y][cur.x] = '.'
			g[ty][tx] = '@'
			return Point{tx, ty}
		}
		return cur
	}
	if moveBox2RL(g, dir, Point{tx, ty}, move, H, W) {
		g[cur.y][cur.x] = '.'
		g[ty][tx] = '@'
		return Point{tx, ty}
	}
	return cur
}

func Day15() {
	inputFile := utils.OpenFile("2024/day15/day15.in")
	outputFile := utils.CreateFile("2024/day15/day15.out")
	defer inputFile.Close()
	defer outputFile.Close()
	reader := bufio.NewReader(inputFile)
	writer := bufio.NewWriter(outputFile)

	var g [][]byte
	var S []byte
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
		tmp := []byte(line)
		g = append(g, tmp)
	}
	for {
		line, err := utils.ReadLine(reader)
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			panic(err)
		}
		tmp := []byte(line)
		S = append(S, tmp...)
	}
	H := len(g)
	W := len(g[0])
	var start Point
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if g[i][j] == '@' {
				start = Point{j, i}
			}
		}
	}
	dir := make(map[string]Point)
	dir["<"] = Point{-1, 0}
	dir[">"] = Point{1, 0}
	dir["^"] = Point{0, -1}
	dir["v"] = Point{0, 1}

	// part 1
	res := 0
	cur := start
	graph1 := make([][]byte, H)
	for i := 0; i < H; i++ {
		graph1[i] = make([]byte, W)
		for j := 0; j < W; j++ {
			graph1[i][j] = g[i][j]
		}
	}
	for i := 0; i < len(S); i++ {
		cur = moveRobot(graph1, dir, cur, S[i], H, W)
	}
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if graph1[i][j] == 'O' {
				res += 100 * i + j
			}
		}
	}
	utils.WriteLine(writer, strconv.Itoa(res))

	// part 2
	res2 := 0
	var cur2 Point
	graph2 := make([][]byte, H)
	for i := 0; i < H; i++ {
		graph2[i] = make([]byte, W * 2)
		for j := 0; j < W; j++ {
			if g[i][j] == '#' {
				graph2[i][j * 2] = '#'
				graph2[i][j * 2 + 1] = '#'
			} else if g[i][j] == 'O' {
				graph2[i][j * 2] = '['
				graph2[i][j * 2 + 1] = ']'
			} else if g[i][j] == '.' {
				graph2[i][j * 2] = '.'
				graph2[i][j * 2 + 1] = '.'
			} else if g[i][j] == '@' {
				graph2[i][j * 2] = '@'
				graph2[i][j * 2 + 1] = '.'
				cur2 = Point{j * 2, i}
			}
		}
	}
	for i := 0; i < len(S); i++ {
		cur2 = moveRobot2(graph2, dir, cur2, S[i], H, W * 2)
	}
	for i := 0; i < H; i++ {
		for j := 0; j < W * 2; j++ {
			if graph2[i][j] == '[' {
				res2 += 100 * i + j
			}
		}
	}
	utils.WriteLine(writer, strconv.Itoa(res2))

	writer.Flush()
}