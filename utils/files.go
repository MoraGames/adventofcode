package utils

import (
	"fmt"
	"os"
)

var (
	FilePathYearPrefix = "y"
	FilePathDayPrefix  = "/day"
	FilePathTestSuffix = "-test"
	FilePathExtension  = ".txt"
)

func ReadFile(year, day int) string {
	file, err := os.ReadFile(fmt.Sprintf("%s%04d%s%02d%s", FilePathYearPrefix, year, FilePathDayPrefix, day, FilePathExtension))
	if err != nil {
		panic(err)
	}
	return string(file)
}

func ReadTestFile(year, day int) string {
	file, err := os.ReadFile(fmt.Sprintf("%s%04d%s%02d%s%s", FilePathYearPrefix, year, FilePathDayPrefix, day, FilePathTestSuffix, FilePathExtension))
	if err != nil {
		panic(err)
	}
	return string(file)
}

func WriteFile(path string, content string) {
	err := os.WriteFile(path, []byte(content), 0644)
	if err != nil {
		panic(err)
	}
}
