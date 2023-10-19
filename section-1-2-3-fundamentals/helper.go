package main

import (
	"os"
)

func sum(a int, b int) int {
	return a + b
}

func saveToFile(fileName string, data []byte) error {
	return os.WriteFile(fileName, data, 0666)

}

func readFromFile(fileName string) ([]byte, error) {
	file, err := os.ReadFile(fileName)

	if err != nil {
		return file, err
	}

	return file, err
}
