package grid

type Fetcher interface {
	FetchGrid() (error, []*City)
}

type Updater interface {
	UpdateGrid(cities []*City) error
}
