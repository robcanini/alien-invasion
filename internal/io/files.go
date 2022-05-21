package io

import (
	"bytes"
	"encoding/csv"
	"os"
)

type FileEntry struct {
	data []string
}

// ReadFile todo: not considering file size. custom encoder to read chunks
func ReadFile(path string) (error, []*FileEntry) {
	err, r := createFileReader(path)
	if err != nil {
		return err, nil
	}
	records, err := r.ReadAll()
	if err != nil {
		return err, nil
	}
	return toFileEntries(records)
}

func toFileEntries(records [][]string) (error, []*FileEntry) {
	entries := make([]*FileEntry, len(records))
	for index, entry := range records {
		entries[index] = &FileEntry{data: entry}
	}
	return nil, entries
}

func createFileReader(path string) (error, *csv.Reader) {
	content, err := os.ReadFile(path)
	r := csv.NewReader(bytes.NewReader(content))
	r.Comma = ' '
	return err, r
}
