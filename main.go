package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	sum, err := parseInputFile("values.txt")
	if err != nil {
		fmt.Println("an error occurred while parsing the input file")
	}

	fmt.Printf("The sum of all of the mul() values is: %d\n", sum)
}

func parseInputFile(filepath string) (sum int, err error) {
	file, err := os.Open(filepath)
	if err != nil {
		return 0, fmt.Errorf("an error occurred while parsing the file {%s}: %v", filepath, err)
	}
	defer file.Close()

	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		matches := re.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			num1, err1 := strconv.Atoi(match[1])
			num2, err2 := strconv.Atoi(match[2])

			if err1 != nil || err2 != nil {
				fmt.Printf("error converting numbers: %v %v\n", err1, err2)
				continue
			}

			sum += num1 * num2
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("an error occurred while reading the file: %v", err)
	}

	return sum, nil
}
