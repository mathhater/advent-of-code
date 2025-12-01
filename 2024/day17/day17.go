package day17

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/mathhater/advent-of-code/utils"
)

const (
	ADV = 0
	BXL = 1
	BST = 2
	JNZ = 3
	BXC = 4
	OUT = 5
	BDV = 6
	CDV = 7
)

func parseRegister(reader *bufio.Reader) int {
	// Register A: 729
	line, err := utils.ReadLine(reader)
	if err != nil {
		panic(err)
	}
	splittedString := strings.Split(line, " ")
	return utils.ParseInt(splittedString[2])
}

func parseProgram(reader *bufio.Reader) []int {
	// Program: 0,1,5,4,3,0
	line, err := utils.ReadLine(reader)
	if err != nil {
		panic(err)
	}
	splittedString := strings.Split(strings.Split(line, " ")[1], ",")
	program := make([]int, len(splittedString))
	for i := 0; i < len(splittedString); i++ {
		program[i] = utils.ParseInt(splittedString[i])
	}
	return program
}

func getValue(operand, A, B, C int) int {
	switch operand {
		case 0: {
			return operand
		}
		case 1: {
			return operand
		}
		case 2: {
			return operand
		}
		case 3: {
			return operand
		}
		case 4: {
			return A
		}
		case 5: {
			return B
		}
		case 6: {
			return C
		}
		case 7: {
			panic("unknown operand")
		}
	}
	panic("unknown operand")
}

func getOutput(programs []int, A, B, C int) []string {
	cur := 0
	var res []string
	for cur < len(programs) {
		switch programs[cur] {
			case ADV: {
				A = A / (1 << getValue(programs[cur + 1], A, B, C))
				cur += 2
				break
			}
			case BXL: {
				B ^= programs[cur + 1]
				cur += 2
				break
			}
			case BST: {
				B = getValue(programs[cur + 1], A, B, C) % 8
				cur += 2
				break
			}
			case JNZ: {
				if A == 0 {
					cur += 2
				} else {
					cur = programs[cur + 1]
				}
				break
			}
			case BXC: {
				B ^= C
				cur += 2
				break
			}
			case OUT: {
				res = append(res, strconv.Itoa(getValue(programs[cur + 1], A, B, C) % 8))
				cur += 2
				break
			}
			case BDV: {
				B = A / (1 << getValue(programs[cur + 1], A, B, C))
				cur += 2
				break
			}
			case CDV:
				C = A / (1 << getValue(programs[cur + 1], A, B, C))
				cur += 2
		}
	}
	return res
}

func isSame(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func Day17() {
	inputFile := utils.OpenFile("2024/day17/day17.in")
	outputFile := utils.CreateFile("2024/day17/day17.out")
	defer inputFile.Close()
	defer outputFile.Close()
	reader := bufio.NewReader(inputFile)
	writer := bufio.NewWriter(outputFile)

	A := parseRegister(reader)
	B := parseRegister(reader)
	C := parseRegister(reader)
	utils.ReadLine(reader)
	programs := parseProgram(reader)

	// part 1
	res := getOutput(programs, A, B, C)
	utils.WriteLine(writer, strings.Join(res, ","))

	// part 2
	// var stringPrograms []string
	// for i := 0; i < len(programs); i++ {
	// 	stringPrograms = append(stringPrograms, strconv.Itoa(programs[i]))
	// }
	// a := 0
	// for {
	// 	tmp := getOutput(programs, a, B, C)
	// 	if isSame(stringPrograms, tmp) && a != A {
	// 		break
	// 	}
	// }
	// utils.WriteLine(writer, strconv.Itoa(a))
	// writer.Flush()
}
