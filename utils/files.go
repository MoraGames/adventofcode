package utils

import (
	"fmt"
	"os"
)

const (
	defaultFilePathFormat           = "./y%04d/day%02d.txt"
	defaultTestFilePathFormat       = "./y%04d/day%02d-test.txt"
	defaultCustomTestFilePathFormat = "./y%04d/day%02d-customtest.txt"
	defaultOutputFilePathFormat     = "./y%04d/outputs/day%02d-%s"
)

var (
	FilePathFormat           = defaultFilePathFormat
	TestFilePathFormat       = defaultTestFilePathFormat
	CustomTestFilePathFormat = defaultCustomTestFilePathFormat
	OutputFilePathFormat     = defaultOutputFilePathFormat
)

func Read(path string, args ...any) string {
	file, err := os.ReadFile(fmt.Sprintf(path, args...))
	if err != nil {
		panic(err)
	}
	return string(file)
}

func ReadFile(year, day int) string {
	return Read(defaultFilePathFormat, year, day)
}

func ReadTestFile(year, day int) string {
	return Read(defaultTestFilePathFormat, year, day)
}

func ReadCustomTestFile(year, day int) string {
	return Read(defaultCustomTestFilePathFormat, year, day)
}

func Write(content string, path string, args ...any) {
	err := os.WriteFile(fmt.Sprintf(path, args...), []byte(content), 0644)
	if err != nil {
		panic(err)
	}
}

func WriteOutput(content string, year, day int, filename string) {
	Write(content, defaultOutputFilePathFormat, year, day, filename)
}
