package main

import (
	"fmt"
	"github.com/robcanini/alien-invasion/internal/io"
)

func main() {
	fmt.Println("Planet X invasion started")

	err, _ := io.ReadFile("")
	if err != nil {
		return
	}
}
