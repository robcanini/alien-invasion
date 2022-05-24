package invasion

import (
	"errors"
	"fmt"
	"github.com/robcanini/alien-invasion/internal/aliens"
	"github.com/robcanini/alien-invasion/internal/grid"
	"github.com/robcanini/alien-invasion/internal/io"
	"sync"
)

type DataSource string

const (
	FileSource DataSource = "file"
)

type Spec struct {
	PlanetName           string
	PlanetGridSourceType DataSource
	PlanetGridSourceUri  string
	AliensNumber         int
}

func Run(spec Spec) error {
	fmt.Printf("Planet %s invasion started\n", spec.PlanetName)
	err, fetcher := buildGridFetcherStrategy(spec.PlanetGridSourceType, spec.PlanetGridSourceUri)
	if err != nil {
		return err
	}
	err = grid.Load(fetcher)
	if err != nil {
		return err
	}
	err, spawnedAliens := aliens.SpreadOn(grid.GetGrid(), spec.AliensNumber)
	if err != nil {
		return err
	}
	startInvasion(spawnedAliens)
	return nil
}

func startInvasion(aliens []*aliens.Alien) {
	var wg sync.WaitGroup
	for _, alien := range aliens {
		wg.Add(1)
		go alien.Startup(&wg)
	}
	wg.Wait()
	grid.PrintGrid()
}

func buildGridFetcherStrategy(sourceType DataSource, sourceUri string) (error, grid.Fetcher) {
	switch sourceType {
	case FileSource:
		return nil, &io.FileFetcher{FilePath: sourceUri}
	}
	return errors.New(fmt.Sprintf("grid fetcher strategy %s has not been implemented", sourceType)), nil
}
