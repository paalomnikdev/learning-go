package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"time"
)

type FileManager struct {
	InputFilePath string
	OutputFilePath string
}

func (fm FileManager) ReadLines() ([]string, error) {
	var lines []string
	file, err := os.Open(fm.InputFilePath)
	
	if err != nil {
		return nil, errors.New("can't open file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		return nil, errors.New("failed to read line in file")
	}

	return lines, nil
}

func (fm FileManager) WriteResult(data any) error {
	file, err := os.Create(fm.OutputFilePath)

	if err != nil {
		return errors.New("failed to create file")
	}

	defer file.Close()

	time.Sleep(3 * time.Second) // simulate slow execution

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		return errors.New("failed to convert data to JSON")
	}

	return nil
}

func New(inputPath, outputPath string) *FileManager {
	return &FileManager{
		InputFilePath: inputPath,
		OutputFilePath: outputPath,
	}
}
