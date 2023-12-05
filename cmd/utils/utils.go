package utils

import (
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
