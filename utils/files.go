package utils

import (
	"fmt"
	"os"
)

func ReadFile(year, day int) string {
	file, err := os.ReadFile("y" + fmt.Sprint(year) + "/day" + fmt.Sprint(day) + ".txt")
	if err != nil {
		panic(err)
	}
	return string(file)
}

func ReadTestFile(year, day int) string {
	file, err := os.ReadFile("y" + fmt.Sprint(year) + "/day" + fmt.Sprint(day) + "-test.txt")
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
