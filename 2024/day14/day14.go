package day14

import (
	"bufio"
	"image"
	"image/color"
	"image/png"
	"os"
	"strconv"
	"strings"

	"github.com/mathhater/advent-of-code/utils"
)

type Point struct {
	x int
	y int
}

type Vector struct {
	x int
	y int
}

type Robot struct {
	p Point
	v Vector
}

func parsePoint(s string) Point {
	splittedStrings := strings.Split(s, ",")
	x := utils.ParseInt(splittedStrings[0][2:])
	y := utils.ParseInt(splittedStrings[1])
	return Point{x, y}
}

func parseVector(s string) Vector {
	splittedStrings := strings.Split(s, ",")
	x := utils.ParseInt(splittedStrings[0][2:])
	y := utils.ParseInt(splittedStrings[1])
	return Vector{x, y}
}

func getCoordinates(originalRobots []Robot, H int, W int) []int {
	robots := make([]Robot, len(originalRobots))
	copy(robots, originalRobots)
	for i := 0; i < 100; i++ {
		for j := 0; j < len(robots); j++ {
			robots[j].p.x += robots[j].v.x
			robots[j].p.y += robots[j].v.y
			if robots[j].p.x < 0 {
				robots[j].p.x += W
			}
			if robots[j].p.x >= W {
				robots[j].p.x %= W
			}
			if robots[j].p.y < 0 {
				robots[j].p.y += H
			}
			if robots[j].p.y >= H {
				robots[j].p.y %= H
			}
		}
	}
	coordinates := make([]int, 4)
	for i := 0; i < len(robots); i++ {
		if robots[i].p.x < W / 2 && robots[i].p.y < H / 2 {
			coordinates[0]++
		} else if robots[i].p.x > W / 2 && robots[i].p.y < H / 2 {
			coordinates[1]++
		} else if robots[i].p.x > W / 2 && robots[i].p.y > H / 2 {
			coordinates[2]++
		} else if robots[i].p.x < W / 2 && robots[i].p.y > H / 2 {
			coordinates[3]++
		}
	}
	return coordinates
}

func makeImageFiles(originalRobots []Robot, H int, W int, count int, flag bool) {
	robots := make([]Robot, len(originalRobots))
	copy(robots, originalRobots)
	for i := 0; i < count; i++ {
		for j := 0; j < len(robots); j++ {
			robots[j].p.x += robots[j].v.x
			robots[j].p.y += robots[j].v.y
			if robots[j].p.x < 0 {
				robots[j].p.x += W
			}
			if robots[j].p.x >= W {
				robots[j].p.x %= W
			}
			if robots[j].p.y < 0 {
				robots[j].p.y += H
			}
			if robots[j].p.y >= H {
				robots[j].p.y %= H
			}
		}
		place := make([][]int, H)
		for j := 0; j < H; j++ {
			place[j] = make([]int, W)
		}
		for j := 0; j < len(robots); j++ {
			place[robots[j].p.y][robots[j].p.x]++
		}

		if flag {
			img := image.NewRGBA(image.Rect(0, 0, W, H))
			white := color.RGBA{255, 255, 255, 255}
			black := color.RGBA{0, 0, 0, 255}
			for j := 0; j < H; j++ {
				for k := 0; k < W; k++ {
					if place[j][k] == 0 {
						img.Set(k, j, white)
					} else {
						img.Set(k, j, black)
					}
				}
			}
			file, err := os.Create("2024/day14/" + strconv.Itoa(i) + ".png")
			if err != nil {
				panic(err)
			}
			defer file.Close()
			err = png.Encode(file, img)
			if err != nil {
				panic(err)
			}
		}
	}
}

func Day14() {
	inputFile := utils.OpenFile("2024/day14/day14.in")
	outputFile := utils.CreateFile("2024/day14/day14.out")
	defer inputFile.Close()
	defer outputFile.Close()
	reader := bufio.NewReader(inputFile)
	writer := bufio.NewWriter(outputFile)

	var robots []Robot
	for {
		line, err := utils.ReadLine(reader)
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			panic(err)
		}

		splittedLine := strings.Split(line, " ")
		p := parsePoint(splittedLine[0])
		v := parseVector(splittedLine[1])
		robots = append(robots, Robot{p, v})
	}
	H := 7
	W := 11

	// part 1
	res := 1
	coordinates := getCoordinates(robots, H, W)
	for i := 0; i < len(coordinates); i++ {
		res *= coordinates[i]
	}
	utils.WriteLine(writer, strconv.Itoa(res))

	// part 2
	// If you actually want to create an image, change the flag to true
	makeImageFiles(robots, H, W, 10000, false)

	writer.Flush()
}