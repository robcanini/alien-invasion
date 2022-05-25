package test

import (
	"github.com/robcanini/alien-invasion/internal/invasion"
	"testing"
)

func TestShouldFileForWrongPlanetGridSourceUriRun(t *testing.T) {
	err := invasion.Run(invasion.Spec{
		PlanetName:           "X",
		PlanetGridSourceType: invasion.FileSource,
		PlanetGridSourceUri:  "not_existing_file.txt",
		AliensNumber:         1,
	})

	if err == nil {
		t.Fatal("engine must fail for spec with a non existing grid source file")
	}
}

func TestShouldFileForWrongInvalidAliensNumberRun(t *testing.T) {
	err := invasion.Run(invasion.Spec{
		PlanetName:           "X",
		PlanetGridSourceType: invasion.FileSource,
		PlanetGridSourceUri:  "not_existing_file.txt",
		AliensNumber:         -1,
	})

	if err == nil {
		t.Fatal("engine must fail for spec with an invalid aliens number")
	}
}

func TestShouldOneAlienRun(t *testing.T) {
	err := invasion.Run(invasion.Spec{
		PlanetName:           "X",
		PlanetGridSourceType: invasion.FileSource,
		PlanetGridSourceUri:  "resource/io_read_file_01.txt",
		AliensNumber:         1,
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestShouldFiveAlienRun(t *testing.T) {
	err := invasion.Run(invasion.Spec{
		PlanetName:           "X",
		PlanetGridSourceType: invasion.FileSource,
		PlanetGridSourceUri:  "resource/io_read_file_01.txt",
		AliensNumber:         5,
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestShouldFailTenAlienExceedCitiesNumberRun(t *testing.T) {
	err := invasion.Run(invasion.Spec{
		PlanetName:           "X",
		PlanetGridSourceType: invasion.FileSource,
		PlanetGridSourceUri:  "resource/io_read_file_01.txt",
		AliensNumber:         10,
	})
	if err == nil {
		t.Fatal(err)
	}
}
