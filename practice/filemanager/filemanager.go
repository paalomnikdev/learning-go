package filemanager

import (
	"bufio"
	"errors"
	"os"
	"encoding/json"
)

func ReadLines(fileName string) ([]string, error) {
	var lines []string
	file, err := os.Open(fileName)
	
	if err != nil {
		return nil, errors.New("can't open file")
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		file.Close()
		return nil, errors.New("failed to read line in file")
	}

	file.Close()

	return lines, nil
}

func WriteJSON(path string, data any) error {
	file, err := os.Create(path)

	if err != nil {
		return errors.New("failed to create file")
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		file.Close()
		return errors.New("failed to convert data to JSON")
	}

	file.Close()

	return nil
}
