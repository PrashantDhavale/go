package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"time"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)
	if err != nil {
		return nil, errors.New("Failed to open file.")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		return nil, errors.New("Reading the file content failed.")
	}

	return lines, nil
}

func (fm FileManager) WriteResults(data any) error {
	file, err := os.Create(fm.OutputFilePath)
	if err != nil {
		return errors.New("Failed to create file.")
	}

	defer file.Close()
	time.Sleep(3 * time.Second)

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		return errors.New("Failed to create file.")
	}

	return nil
}

func New(inputPath, outputPath string) FileManager {
	return FileManager{
		InputFilePath:  inputPath,
		OutputFilePath: outputPath,
	}
}
