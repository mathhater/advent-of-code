package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func openFile(fileName string) *os.File {
	inputFile, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	return inputFile
}

func createFile(fileName string) *os.File {
	outputFile, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	return outputFile
}

func readLine(reader *bufio.Reader) (string, error) {
	line, err := reader.ReadString('\n')
	return strings.Replace(line, "\r\n", "", -1), err
}

func writeLine(writer *bufio.Writer, str string) {
	_, err := writer.WriteString(str + "\n")
	if err != nil {
		panic(err)
	}
}

func getInt(str string) int {
	res, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		panic(err)
	}
	return int(res)
}

func main() {
	inputFile := openFile("day1.in")
	outputFile := createFile("day_1.out")
	defer inputFile.Close()
	defer outputFile.Close()
	reader := bufio.NewReader(inputFile)
	writer := bufio.NewWriter(outputFile)

	var A, B []int
	for {
		line, err := readLine(reader)
		if err.Error() == "EOF" {
			break
		}
		if err != nil {
			panic(err)
		}

		splittedLine := strings.Split(line, "   ")
		a := getInt(splittedLine[0])
		b := getInt(splittedLine[1])
		A = append(A, int(a))
		B = append(B, int(b))
	}
	sort.Ints(A)
	sort.Ints(B)

	// part 1
	res := 0
	for i := 0; i < len(A); i++ {
		res += abs(B[i] - A[i])
	}
	writeLine(writer, strconv.Itoa(res))

	// part 2
	res2 := 0
	var cnt = make(map[int]int)
	for i := 0; i < len(A); i++ {
		cnt[B[i]]++
	}
	for i := 0; i < len(A); i++ {
		res2 += A[i] * cnt[A[i]]
	}
	writeLine(writer, strconv.Itoa(res2))
	writer.Flush()
}