package shared

import (
	"bufio"
	"os"
	"path/filepath"
)

func ReadFile(filename string) ([]string, error) {
	var lines []string

	absPath, _ := filepath.Abs(filename)
	file, err := os.Open(absPath)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
