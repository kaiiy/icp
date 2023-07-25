package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

const Version = "v0.3.0"

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	if os.Args[1] == "-v" || os.Args[1] == "--version" {
		fmt.Println("icp:", Version)
		os.Exit(0)
	}

	inputFile := os.Args[1]
	number, suffix, err := parseFileName(inputFile)
	if err != nil {
		fmt.Println("Failed to parse file name.")
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

func printUsage() {
	fmt.Println(
`Icp: Incremental File Copy

Commands:
  [filename]     Copy file and increment number within filename
  -v, --version  Print the version`)
}

func parseFileName(inputFile string) (number int, fileName string, err error) {
	regex := regexp.MustCompile(`^(\d+)(_.+\..+)`)
	matches := regex.FindStringSubmatch(inputFile)

	if len(matches) != 3 {
		return 0, "", fmt.Errorf("invalid file name")
	}

	number, err = strconv.Atoi(matches[1])
	if err != nil {
		return 0, "", err
	}

	return number, matches[2], nil
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
