package utils

import "os"

func OpenFile(fileName string) *os.File {
	inputFile, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	return inputFile
}

func CreateFile(fileName string) *os.File {
	outputFile, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	return outputFile
}