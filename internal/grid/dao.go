package grid

type Fetcher interface {
	fetchGrid() (error, []*City)
}

type Updater interface {
	updateGrid(cities []*City) error
}
