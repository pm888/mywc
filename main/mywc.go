package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"unicode"
)

func main() {
	countLines := flag.Bool("l", false, "Count lines")
	countWords := flag.Bool("w", false, "Count words")
	countChars := flag.Bool("c", false, "Count characters")
	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Println("Usage: mywc [options] <filename> | l - lines w - words c - characters  ")
		os.Exit(1)
	}

	filename := os.Args[len(os.Args)-1]

	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lineCount := 0
	wordCount := 0
	charCount := 0

	for scanner.Scan() {
		lineCount++
		charCount += len(scanner.Text()) + 1
		words := splitWords(scanner.Text())
		wordCount += len(words)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	if *countLines {
		fmt.Printf("Lines: %d\n", lineCount)
	}
	if *countWords {
		fmt.Printf("Words: %d\n", wordCount)
	}
	if *countChars {
		fmt.Printf("Characters: %d\n", charCount)
	}
}

func splitWords(text string) []string {
	words := make([]string, 0)
	currentWord := ""
	inWord := false

	for _, char := range text {
		if unicode.IsSpace(char) {
			if inWord {
				words = append(words, currentWord)
				currentWord = ""
				inWord = false
			}
		} else {
			currentWord += string(char)
			inWord = true
		}
	}

	if inWord {
		words = append(words, currentWord)
	}

	return words
}
