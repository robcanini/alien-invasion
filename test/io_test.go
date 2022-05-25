package test

import (
	"github.com/robcanini/alien-invasion/internal/io"
	"testing"
)

func TestShouldBeSevenEntriesReadFile(t *testing.T) {
	err, entries := io.ReadFile("resource/io_read_file_01.txt")
	if err != nil {
		t.Fatal(err)
	}
	if len(entries) != 7 {
		t.Fatal("entries size should equal to 7")
	}
}

func TestShouldFailReadFile(t *testing.T) {
	err, _ := io.ReadFile("not_existing_file.txt")
	if err == nil {
		t.Fatal("readfile should fail for not existing file")
	}
}

func TestShouldBeThreeEntryDataReadFile(t *testing.T) {
	err, entries := io.ReadFile("resource/io_read_file_02.txt")
	if err != nil {
		t.Fatal(err)
	}
	if len(entries) != 1 {
		t.Fatal("expecting to have 1 entry")
	}
	if len((*entries[0]).Data) != 3 {
		t.Fatal("entry data size should equal to 3")
	}
}

func TestShouldNotFailFetchGrid(t *testing.T) {
	fileFetcher := &io.FileFetcher{FilePath: "resource/io_read_file_01.txt"}
	err, _ := fileFetcher.FetchGrid()
	if err != nil {
		t.Fatal(err)
	}
}

func TestShouldFailFetchGrid(t *testing.T) {
	fileFetcher := &io.FileFetcher{FilePath: "not_existing_file.txt"}
	err, _ := fileFetcher.FetchGrid()
	if err == nil {
		t.Fatal("FetchGrid should fail for not existing file")
	}
}

func TestShouldBeSevenParsedCityFetchGrid(t *testing.T) {
	fileFetcher := &io.FileFetcher{FilePath: "resource/io_read_file_01.txt"}
	err, grid := fileFetcher.FetchGrid()
	if err != nil {
		t.Fatal(err)
	}
	if len(grid) != 7 {
		t.Fatal("grid size should equal to 7")
	}
}

func TestShouldBeTwoRoadsParsedFetchGrid(t *testing.T) {
	fileFetcher := &io.FileFetcher{FilePath: "resource/io_read_file_02.txt"}
	err, grid := fileFetcher.FetchGrid()
	if err != nil {
		t.Fatal(err)
	}
	if len(grid) != 1 {
		t.Fatal("expecting to have 1 single entry city")
	}
	if len((*grid[0]).Roads) != 2 {
		t.Fatal("roads slice should be of size 2")
	}
}
