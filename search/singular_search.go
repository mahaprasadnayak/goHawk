package search

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

func SearchFile(filename, pattern string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// If the pattern is empty, match every line.
	var re *regexp.Regexp
	if pattern == "" {
		re = regexp.MustCompile(".*") // Matches any line.
	} else {
		re, err = regexp.Compile(pattern)
		if err != nil {
			return err
		}
	}

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n') // Read line by line
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		// Trim the newline character and check if the line matches the pattern
		line = line[:len(line)-1]
		if re.MatchString(line) {
			fmt.Printf("%s: %s\n", filename, line)
		}
	}
	return nil
}
