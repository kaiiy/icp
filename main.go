package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: icp [file]")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	number, suffix, err := parseFileName(inputFile)
	if err != nil {
		fmt.Println("Failed to parse the numeric part of the file name.")
		os.Exit(1)
	}

	outputFile := strconv.Itoa(number+1) + suffix
	err = copyFile(inputFile, outputFile)
	if err != nil {
		fmt.Println("Failed to copy file.")
		os.Exit(1)
	}

	fmt.Printf("File copied to: %s\n", outputFile)
}

func parseFileName(inputFile string) (number int, suffix string, err error) {
	baseName := filepath.Base(inputFile)
	nonNumericPart := strings.TrimLeftFunc(baseName, func(r rune) bool { return '0' <= r && r <= '9' })
	prefix := strings.TrimSuffix(baseName, nonNumericPart)

	number, err = strconv.Atoi(prefix)
	if err != nil {
		return 0, "", err
	}

	return number, nonNumericPart, nil
}

func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return err
	}

	return nil
}
