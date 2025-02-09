package search

import (
	"log"
	"os"
	"path/filepath"
	"regexp"
)

func SearchDirectory(directory, pattern string ,caseInsensitive,showLineNumbers,countMatches bool) error {
	
	if caseInsensitive {
		pattern = "(?i)" + pattern
	}

	re, compErr := regexp.Compile(pattern)
	if compErr != nil {
		return compErr
	}
	/*
	alt method to implement regex
	// Compile the regex 
	// var re *regexp.Regexp
	// if pattern == "" {
	// 	re = regexp.MustCompile(".*") // Matches any line.
	// } else {
	// 	var err error
	// 	re, err = regexp.Compile(pattern)
	// 	if err != nil {
	// 		return err
	// 	}
	// }
	*/
	// Walk the directory tree
	err := filepath.WalkDir(directory, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			log.Printf("Error accessing path %s: %s\n", path, err.Error())
			return nil
		}

		// Skip directories
		if d.IsDir() {
			return nil
		}

		// Search the file
		return SearchFile(path, re.String(),caseInsensitive ,showLineNumbers,countMatches)
	})

	return err
}
