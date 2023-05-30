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
		usageAndExit()
	}

	inputFile := os.Args[1]
	number, suffix, err := parseFileName(inputFile)
	if err != nil {
		exitWithError("Failed to parse the numeric part of the file name.")
	}

	outputFile := strconv.Itoa(number+1) + suffix
	if err := copyFile(inputFile, outputFile); err != nil {
		exitWithError("Failed to copy file.")
	}

	fmt.Printf("File copied to: %s\n", outputFile)
}

func usageAndExit() {
	fmt.Println("Usage: icp [file]")
	os.Exit(1)
}

func exitWithError(errMsg string) {
	fmt.Println(errMsg)
	os.Exit(1)
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
