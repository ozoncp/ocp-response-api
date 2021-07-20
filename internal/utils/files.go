package utils

import (
	"bytes"
	"os"
)

// RepeatableRead reads file item times and return file content
func RepeatableRead(filePath string, iter uint64) ([]string, error) {

	readFile := func() (content string, err error) {
		file, err := os.OpenFile(filePath, os.O_RDONLY, 0666)
		if err != nil {
			return "", err
		}

		defer file.Close()

		data := new(bytes.Buffer)
		if _, err = data.ReadFrom(file); err != nil {
			return "", nil
		}

		return data.String(), nil
	}

	result := make([]string, 0, iter)

	var i uint64
	for i = 0; i < iter; i++ {
		if data, err := readFile(); err != nil {
			return nil, err
		} else {
			result = append(result, data)
		}
	}

	return result, nil
}
