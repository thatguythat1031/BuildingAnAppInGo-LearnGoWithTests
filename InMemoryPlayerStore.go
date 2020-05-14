package main

//InMemoryPlayerStore stores player scores in memory
type InMemoryPlayerStore struct {
	store map[string]int
}

func (i *InMemoryPlayerStore) GetLeague() []Player {
	var league []Player
	for name, wins := range i.store {
		league = append(league, Player{name, wins})
	}
	return league
}

//NewInMemoryPlayerStore creates a new InMemoryPlayerStore
func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

//RecordWin adds a point to the player's score
func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}

//GetPlayerScore retrieves the given player's score
func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}
