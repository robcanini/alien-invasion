package aliens

func Generate(number int) []*Alien {
	aliens := make([]*Alien, number)
	for index := 0; index < number; index++ {
		aliens[index] = Create()
	}
	return aliens
}
