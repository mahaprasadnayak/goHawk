package search

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
)

func SearchFile(filename, pattern string,caseInsensitive,showLineNumbers bool) error {
	if len(pattern) > 0 && len(pattern) < 3 {
		return errors.New("pattern must contain at least 3 characters")
	}
	file, err := os.Open(filename)
	if err != nil {
		return errors.New("failed to open the file: " + filename)
	}
	defer file.Close()
	if caseInsensitive {
		pattern = "(?i)" + pattern
	}
	/*
	alt method to implement regex
	// If the pattern is empty, match every line.
	//var re *regexp.Regexp
	// if pattern == "" {
	// 	re = regexp.MustCompile(".*") // Matches any line.
	// } else {
	// 	re, err = regexp.Compile("(?i).*" + regexp.QuoteMeta(pattern) + ".*")
	// 	if err != nil {
	// 		return err
	// 	}
	// }
	*/
	re, err := regexp.Compile(pattern)
	if err != nil {
		return err
	}
	reader := bufio.NewScanner(file)
	lineNumber:=0
	for reader.Scan() {
		lineNumber++
		line := reader.Text()
		if re.MatchString(line) {
			if showLineNumbers {
				fmt.Printf("%s: line number :: %d: %s\n", filename, lineNumber, line)
			}else{
				fmt.Printf("%s: %s\n", filename, line)
			}
		}
	}
	if err := reader.Err(); err != nil {
		return errors.New("failed to finish reading the file: " + filename)
	}
	return nil
}
