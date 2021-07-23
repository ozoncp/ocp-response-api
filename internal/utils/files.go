package utils

import (
	"bytes"
	"errors"
	"os"
	"time"
)

// RepeatableRead reads file item times and return file content
func RepeatableRead(filePath string, iter int) ([]string, error) {
	if iter <= 0 {
		return nil, errors.New("iter must be greater than 0")
	}

	readFile := func() (string, error) {
		file, err := os.OpenFile(filePath, os.O_RDONLY, 0666)
		if err != nil {
			return "", err
		}

		defer file.Close()

		data := new(bytes.Buffer)
		if _, err = data.ReadFrom(file); err != nil {
			return "", err
		}

		return data.String(), nil
	}

	result := make([]string, 0, iter)

	var fileOpenError error
	for i := 0; i < iter; i++ {
		if data, err := readFile(); err != nil {
			fileOpenError = err
			time.Sleep(1 * time.Second)
		} else {
			result = append(result, data)
			fileOpenError = nil
		}
	}

	if len(result) == 0 {
		return nil, fileOpenError
	}

	return result, nil
}
