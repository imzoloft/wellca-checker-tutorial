package io

import (
	"errors"
	"os"
	"strings"
	"sync"
)

var fileMutex sync.Mutex

func ReadFile(fileName string) ([]string, error) {
	file, err := os.ReadFile(fileName)
	if err != nil {
		return nil, errors.New("unable to read file")
	}

	lines := strings.Split(string(file), "\n")

	return lines, nil
}

func WriteToFile(fileName string, message string) error {
	fileMutex.Lock()
	defer fileMutex.Unlock()

	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := file.WriteString(message + "\n"); err != nil {
		return err
	}
	return nil
}
