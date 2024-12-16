package day16

import (
	"bufio"
	"container/heap"
	"math"
	"strconv"

	"github.com/mathhater/advent-of-code/utils"
)

type Point struct {
	x, y int
}

type Node struct {
	p Point
	w, d int
	path []Point
}

type PriorityQueue []Node

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].w < pq[j].w }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Top() Node { return (*pq)[0] }

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(Node))
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}

func getScore (g [][]byte, start Point, H, W int) int {
	res := math.MaxInt32
	dir := [4]Point{{1, 0}, {0, -1}, {-1, 0}, {0, 1}}
	pq := &PriorityQueue{}
	dp := make([][][]int, H)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([][]int, W)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = make([]int, 4)
			for k := 0; k < 4; k++ {
				dp[i][j][k] = math.MaxInt32
			}
		}
	}
	heap.Push(pq, Node{start, 0, 0, nil})
	dp[start.y][start.x][0] = 0
	for i := 1; i < len(dir); i++ {
		heap.Push(pq, Node{start, 1000, i, nil})
		dp[start.y][start.x][i] = 1000
	}
	for pq.Len() > 0 {
		n := pq.Top()
		heap.Pop(pq)
		p := n.p
		if g[p.y][p.x] == 'E' && res > n.w {
			res = n.w
			continue
		}
		td := n.d
		ty := p.y + dir[td].y
		tx := p.x + dir[td].x
		if !(ty < 0 || ty >= H || tx < 0 || tx >= W || dp[ty][tx][td] <= n.w + 1|| g[ty][tx] == '#') {
			dp[ty][tx][td] = n.w + 1
			heap.Push(pq, Node{Point{tx, ty}, n.w + 1, td, n.path})
		}
		td = (n.d + 1) % 4
		ty = p.y
		tx = p.x
		if !(ty < 0 || ty >= H || tx < 0 || tx >= W || dp[ty][tx][td] <= n.w + 1000 || g[ty][tx] == '#') {
			dp[ty][tx][td] = n.w + 1000
			heap.Push(pq, Node{Point{tx, ty}, n.w + 1000, td, n.path})
		}
		td = (n.d - 1 + 4) % 4
		ty = p.y
		tx = p.x
		if !(ty < 0 || ty >= H || tx < 0 || tx >= W || dp[ty][tx][td] <= n.w + 1000 || g[ty][tx] == '#') {
			dp[ty][tx][td] = n.w + 1000
			heap.Push(pq, Node{Point{tx, ty}, n.w + 1000, td, n.path})
		}
	}
	return res
}

func getBestPaths (g [][]byte, start Point, H, W int) int {
	dir := [4]Point{{1, 0}, {0, -1}, {-1, 0}, {0, 1}}
	pq := &PriorityQueue{}
	dp := make([][][]int, H)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([][]int, W)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = make([]int, 4)
			for k := 0; k < 4; k++ {
				dp[i][j][k] = math.MaxInt32
			}
		}
	}
	heap.Push(pq, Node{start, 0, 0, []Point{start}})
	dp[start.y][start.x][0] = 0
	for i := 1; i < len(dir); i++ {
		heap.Push(pq, Node{start, 1000, i, []Point{start}})
		dp[start.y][start.x][i] = 1000
	}
	for pq.Len() > 0 {
		n := pq.Top()
		heap.Pop(pq)
		p := n.p
		if g[p.y][p.x] == 'E' {
			pathMap := make(map[Point]bool)
			for _, np := range n.path {
				pathMap[np] = true
			}
			for pq.Len() > 0 {
				x := pq.Top()
				heap.Pop(pq)
				if x.p.y == p.y && x.p.x == p.x && x.w == n.w {
					for _, xp := range x.path {
						pathMap[xp] = true
					}
				}
			}
			for key, _ := range pathMap {
				g[key.y][key.x] = 'O'
			}
			return len(pathMap)
		}
		td := n.d
		ty := p.y + dir[td].y
		tx := p.x + dir[td].x
		tpath1 := make([]Point, len(n.path))
		copy(tpath1, n.path)
		tpath1 = append(tpath1, Point{tx, ty})
		if !(ty < 0 || ty >= H || tx < 0 || tx >= W || dp[ty][tx][td] < n.w + 1 || g[ty][tx] == '#') {
			dp[ty][tx][td] = n.w + 1
			heap.Push(pq, Node{Point{tx, ty}, n.w + 1, td, tpath1})
		}
		td = (n.d + 1) % 4
		ty = p.y
		tx = p.x
		tpath2 := make([]Point, len(n.path))
		copy(tpath2, n.path)
		tpath2 = append(tpath2, Point{tx, ty})
		if !(ty < 0 || ty >= H || tx < 0 || tx >= W || dp[ty][tx][td] < n.w + 1000 || g[ty][tx] == '#') {
			dp[ty][tx][td] = n.w + 1000
			heap.Push(pq, Node{Point{tx, ty}, n.w + 1000, td, tpath2})
		}
		td = (n.d - 1 + 4) % 4
		ty = p.y
		tx = p.x
		tpath3 := make([]Point, len(n.path))
		copy(tpath3, n.path)
		tpath3 = append(tpath3, Point{tx, ty})
		if !(ty < 0 || ty >= H || tx < 0 || tx >= W || dp[ty][tx][td] < n.w + 1000 || g[ty][tx] == '#') {
			dp[ty][tx][td] = n.w + 1000
			heap.Push(pq, Node{Point{tx, ty}, n.w + 1000, td, tpath3})
		}
	}
	return 0
}

func Day16() {
	inputFile := utils.OpenFile("2024/day16/day16.in")
	outputFile := utils.CreateFile("2024/day16/day16.out")
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
	var start Point
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if g[i][j] == 'S' {
				start = Point{j, i}
			}
		}
	}

	// part 1
	res := getScore(g, start, H, W)
	utils.WriteLine(writer, strconv.Itoa(res))

	// part 2
	res2 := getBestPaths(g, start, H, W)
	utils.WriteLine(writer, strconv.Itoa(res2))

	writer.Flush()
}