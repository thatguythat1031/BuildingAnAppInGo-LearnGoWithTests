package main

import (
	"log"
	"net/http"
)

//InMemoryPlayerStore stores player scores in memory
type InMemoryPlayerStore struct {
	store map[string]int
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

func main() {
	server := &PlayerServer{NewInMemoryPlayerStore()}

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
