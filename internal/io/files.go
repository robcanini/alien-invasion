package io

import (
	"bufio"
	"os"
	"strings"
)

type FileEntry struct {
	data []string
}

func ReadFile(path string) (error, []*FileEntry) {
	err, scanner := createFileScanner(path)
	if err != nil {
		return err, nil
	}
	entries := make([]*FileEntry, 0)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), " ")
		entries = append(entries, &FileEntry{data: data})
	}
	return nil, entries
}

func createFileScanner(path string) (error, *bufio.Scanner) {
	file, err := os.Open(path)
	if err != nil {
		return err, nil
	}
	return nil, bufio.NewScanner(file)
}
