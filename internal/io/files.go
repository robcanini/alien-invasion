package io

import (
	"encoding/csv"
	"fmt"
	"log"
	"strings"
	"sync"
)

var mu sync.Mutex

type FileEntry struct {
	data []string
}

func readFile(path string) []*FileEntry {

	// todo
	
	in := `first_name;last_name;username
		"Rob";"Pike";rob
		# lines beginning with a # character are ignored
		Ken;Thompson;ken
		"Robert";"Griesemer";"gri"
		`

	r := csv.NewReader(strings.NewReader(in))
	r.Comma = ' '

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(records)

	return nil
}
