package utils

import (
	"bufio"
	"strings"
)

func ReadLine(reader *bufio.Reader) (string, error) {
	line, err := reader.ReadString('\n')
	return strings.Replace(line, "\r\n", "", -1), err
}

func WriteLine(writer *bufio.Writer, str string) {
	_, err := writer.WriteString(str + "\n")
	if err != nil {
		panic(err)
	}
}