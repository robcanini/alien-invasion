package io

import (
	"errors"
	"fmt"
	"github.com/robcanini/alien-invasion/internal/grid"
	"strings"
)

type FileDao struct {
	FilePath string
}

func (fileDao *FileDao) FetchGrid() (error, []*grid.City) {
	err, entries := ReadFile(fileDao.FilePath)
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
		return errors.New("file entry must contain city name and at least one direction"), nil
	}
	cityName := entryData[0]
	err, roads := extractCityRoads(entryData, cityName, refsMap)
	if err != nil {
		return err, nil
	}
	var city = (*refsMap)[cityName]
	if city == nil {
		city = &grid.City{Name: cityName}
		(*refsMap)[cityName] = city
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
			destinationCity = &grid.City{Name: dirData[1]}
			(*refsMap)[dirData[1]] = destinationCity
		}
		roads[index] = &grid.Road{
			Direction:   grid.Direction(dirData[0]),
			Destination: destinationCity,
		}
	}
	return nil, roads
}

func enrichErrorWithIndex(err error, index int) error {
	return errors.New(fmt.Sprintf("parsing error at line %d: %s", index, err.Error()))
}
