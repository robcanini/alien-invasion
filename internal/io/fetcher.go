package io

import (
	"errors"
	"fmt"
	"github.com/robcanini/alien-invasion/internal/grid"
	"strings"
)

type FileFetcher struct {
	FilePath string
}

func (fileFetcher *FileFetcher) FetchGrid() (error, []*grid.City) {
	err, entries := ReadFile(fileFetcher.FilePath)
	if err != nil {
		return err, nil
	}
	err, cities := parseEntries(entries)
	if err != nil {
		return err, nil
	}
	return nil, cities
}

func parseEntries(entries []*FileEntry) (error, []*grid.City) {
	cities := make([]*grid.City, len(entries))
	refsMap := make(map[string]*grid.City)
	for index, entry := range entries {
		err, city := toGridCity(entry, &refsMap)
		if err != nil {
			return enrichErrorWithIndex(err, index), nil
		}
		cities[index] = city
	}
	return nil, cities
}

func toGridCity(entry *FileEntry, refsMap *map[string]*grid.City) (error, *grid.City) {
	entryData := entry.data
	if len(entryData) <= 1 {
		return errors.New("city entry must contain name and at least one road to another city"), nil
	}
	cityName := entryData[0]
	err, roads := extractCityRoads(entryData, cityName, refsMap)
	if err != nil {
		return err, nil
	}
	var city = (*refsMap)[cityName]
	if city == nil {
		city = createCity(cityName, refsMap)
	}
	city.Roads = roads
	return nil, city
}

func extractCityRoads(entryData []string, cityName string, refsMap *map[string]*grid.City) (error, []*grid.Road) {
	roads := make([]*grid.Road, len(entryData)-1)
	for index, road := range entryData[1:] {
		dirData := strings.Split(road, "=")
		if len(dirData) != 2 {
			return errors.New("city roads must be specified in the format 'west=Foo'"), nil
		}
		if dirData[1] == cityName {
			return errors.New("city roads cannot contain a self reference"), nil
		}
		destinationCity := (*refsMap)[dirData[1]]
		if destinationCity == nil {
			destinationCity = createCity(dirData[1], refsMap)
		}
		roads[index] = &grid.Road{
			Direction:   grid.Direction(dirData[0]),
			Destination: destinationCity,
		}
	}
	return nil, roads
}

func createCity(cityName string, refsMap *map[string]*grid.City) *grid.City {
	destinationCity := grid.CreateCity(cityName)
	(*refsMap)[cityName] = destinationCity
	return destinationCity
}

func enrichErrorWithIndex(err error, index int) error {
	return errors.New(fmt.Sprintf("parsing error at line %d: %s", index+1, err.Error()))
}
