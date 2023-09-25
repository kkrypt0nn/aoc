package aoc

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

// ReadFile reads the contents of the file and returns a slice of strings.
func ReadFile(fileName string, testing ...bool) []string {
	if len(testing) > 0 && testing[0] {
		fileName = strings.Replace(fileName, ".txt", "_test.txt", 1)
	}
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	file.Close()
	return text
}

// ReadFileAsInteger reads the contents of the file and returns a slice of integers.
func ReadFileAsInteger(fileName string, testing ...bool) []int {
	if len(testing) > 0 && testing[0] {
		fileName = strings.Replace(fileName, ".txt", "_test.txt", 1)
	}
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []int
	for scanner.Scan() {
		converted, _ := strconv.Atoi(scanner.Text())
		text = append(text, converted)
	}
	file.Close()
	return text
}

// ReadFileAsString reads the contents of the file and returns a string.
func ReadFileAsString(fileName string, testing ...bool) string {
	if len(testing) > 0 && testing[0] {
		fileName = strings.Replace(fileName, ".txt", "_test.txt", 1)
	}
	file, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	return string(file)
}
