package utils

import (
	"bufio"
	"fmt"
	"os"
)

func OpenInputFile(filename string) *os.File {
	file, err := os.Open(fmt.Sprintf("./cmd/inputs/%s", filename))

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return file
}

func ReadInputFile(filename string, fn func(*bufio.Scanner)) {
	file := OpenInputFile(filename)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fn(scanner)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
