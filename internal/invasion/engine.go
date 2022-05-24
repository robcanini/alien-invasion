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

type Result struct {
}

func Run(spec Spec) (error, *Result) {
	fmt.Printf("Planet %s invasion started\n", spec.PlanetName)
	err, fetcher := buildGridFetcherStrategy(spec.PlanetGridSourceType, spec.PlanetGridSourceUri)
	if err != nil {
		return err, nil
	}
	err = grid.Load(fetcher)
	if err != nil {
		return err, nil
	}
	err, spawnedAliens := aliens.SpreadOn(grid.GetGrid(), spec.AliensNumber)
	if err != nil {
		return err, nil
	}
	return startInvasion(spawnedAliens)
}

func startInvasion(aliens []*aliens.Alien) (error, *Result) {
	var wg sync.WaitGroup
	for _, alien := range aliens {
		wg.Add(1)
		go alien.Startup(&wg)
	}
	wg.Wait()
	return nil, extractResult()
}

func extractResult() *Result {
	return nil
}

func buildGridFetcherStrategy(sourceType DataSource, sourceUri string) (error, grid.Fetcher) {
	switch sourceType {
	case FileSource:
		return nil, &io.FileFetcher{FilePath: sourceUri}
	}
	return errors.New(fmt.Sprintf("grid fetcher strategy %s has not been implemented", sourceType)), nil
}
