package day08

import (
	"bufio"
	"strconv"

	"github.com/mathhater/advent-of-code/utils"
)

type Point struct {
	x int
	y int
}

func Day08() {
	inputFile := utils.OpenFile("2024/day08/day8.in")
	outputFile := utils.CreateFile("2024/day08/day8.out")
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
	A := make(map[string][]Point)
	for i := 0; i < H; i++ {
		slice := []byte(g[i])
		for j := 0; j < W; j++ {
			if g[i][j] != '.' {
				ch := string(slice[j])
				if _, exists := A[ch]; !exists {
					A[ch] = make([]Point, 0)
				}
				A[ch] = append(A[ch], Point{x: j, y: i})
			}
		}
	}

	// part 1
	antinodes := make(map[Point]bool, 0)
	for node, coordinates := range A {
		for i := 0; i < len(coordinates); i++ {
			for j := i + 1; j < len(coordinates); j++ {
				p1 := coordinates[i]
				p2 := coordinates[j]
				distX := p2.x - p1.x
				distY := p2.y - p1.y
				p3 := Point{x: p2.x + distX, y: p2.y + distY}
				p4 := Point{x: p2.x - distX, y: p2.y - distY}
				p5 := Point{x: p1.x + distX, y: p1.y + distY}
				p6 := Point{x: p1.x - distX, y: p1.y - distY}
				if 0 <= p3.x && p3.x < W && 0 <= p3.y && p3.y < H && g[p3.y][p3.x] != node[0] {
					antinodes[p3] = true
					p3 = Point{x: p3.x + distX, y: p3.y + distY}
				}
				if 0 <= p4.x && p4.x < W && 0 <= p4.y && p4.y < H && g[p4.y][p4.x] != node[0] {
					antinodes[p4] = true
				}
				if 0 <= p5.x && p5.x < W && 0 <= p5.y && p5.y < H && g[p5.y][p5.x] != node[0] {
					antinodes[p5] = true
				}
				if 0 <= p6.x && p6.x < W && 0 <= p6.y && p6.y < H && g[p6.y][p6.x] != node[0] {
					antinodes[p6] = true
				}
			}
		}
	}
	res := len(antinodes)
	utils.WriteLine(writer, strconv.Itoa(res))

	// part 2
	antinodes2 := make(map[Point]bool, 0)
	for node, coordinates := range A {
		for i := 0; i < len(coordinates); i++ {
			p := coordinates[i]
			antinodes2[p] = true
			for j := i + 1; j < len(coordinates); j++ {
				p1 := coordinates[i]
				p2 := coordinates[j]
				distX := p2.x - p1.x
				distY := p2.y - p1.y
				p3 := Point{x: p2.x + distX, y: p2.y + distY}
				p4 := Point{x: p2.x - distX, y: p2.y - distY}
				p5 := Point{x: p1.x + distX, y: p1.y + distY}
				p6 := Point{x: p1.x - distX, y: p1.y - distY}
				for {
					if !(0 <= p3.x && p3.x < W && 0 <= p3.y && p3.y < H && g[p3.y][p3.x] != node[0]) {
						break
					}
					antinodes2[p3] = true
					p3 = Point{x: p3.x + distX, y: p3.y + distY}
				}
				for {
					if !(0 <= p4.x && p4.x < W && 0 <= p4.y && p4.y < H && g[p4.y][p4.x] != node[0]) {
						break
					}
					antinodes2[p4] = true
					p4 = Point{x: p4.x - distX, y: p4.y - distY}
				}
				for {
					if !(0 <= p5.x && p5.x < W && 0 <= p5.y && p5.y < H && g[p5.y][p5.x] != node[0]) {
						break
					}
					antinodes2[p5] = true
					p5 = Point{x: p5.x + distX, y: p5.y + distY}
				}
				for {
					if !(0 <= p6.x && p6.x < W && 0 <= p6.y && p6.y < H && g[p6.y][p6.x] != node[0]) {
						break
					}
					antinodes2[p6] = true
					p6 = Point{x: p6.x - distX, y: p6.y - distY}
				}
			}
		}
	}
	res2 := len(antinodes2)
	for coordinates := range antinodes2 {
		x := coordinates.x
		y := coordinates.y
		slice := []byte(g[y])
		slice[x] = '#'
		g[y] = string(slice)
	}
	utils.WriteLine(writer, strconv.Itoa(res2))
	writer.Flush()
}