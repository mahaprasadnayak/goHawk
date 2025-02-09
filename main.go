package main

import (
	"fmt"
	"log"

	"github.com/mahaprasadnayak/goHawk/search"
	"github.com/spf13/cobra"
)

func main() {
	var pattern string
	var file string
	var recursive, caseInsensitive,showLineNumbers,countMatches  bool

	// Create the root command
	var rootCmd = &cobra.Command{
		Use:   "goHawk",
		Short: "goHawk - High-performance text search tool",
		Run: func(cmd *cobra.Command, args []string) {
			if recursive {
				if file == "" {
					fmt.Println("Usage: goHawk --pattern <pattern> --file <directory> -r")
					return
				}
				fmt.Printf("Searching recursively for pattern: '%s' in directory: '%s'\n", pattern, file)
				searchErr := search.SearchDirectory(file, pattern,caseInsensitive,showLineNumbers,countMatches)
				if searchErr != nil {
					log.Println("Error in recursive search:", searchErr.Error())
				}
			} else {
				if pattern == "" || file == "" {
					fmt.Println("Usage: goHawk --pattern <pattern> --file <filename>")
					return
				}
				fmt.Printf("Searching for pattern: '%s' in file: '%s'\n", pattern, file)
				searchErr := search.SearchFile(file, pattern,caseInsensitive,showLineNumbers,countMatches)
				if searchErr != nil {
					log.Println("Error in file search:", searchErr.Error())
				}
			}
		},
	}

	// Define flags
	rootCmd.Flags().StringVar(&pattern, "pattern", "", "Pattern to search for")
	rootCmd.Flags().StringVar(&file, "file", "", "File or directory to search in")
	rootCmd.Flags().BoolVarP(&recursive, "recursive", "r", false, "Recursively search directories")
	rootCmd.Flags().BoolVarP(&caseInsensitive, "ignore-case", "i", false, "Perform case-insensitive search")
	rootCmd.Flags().BoolVarP(&showLineNumbers, "line-numbers", "n", false, "Show line numbers")
	rootCmd.Flags().BoolVarP(&countMatches, "count", "c", false, "Show count of matches")


	// Execute the command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
