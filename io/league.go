package main

import (
	"encoding/json"
	"fmt"
	"io"
)

// League is a slice of players
type League []Player

// NewLeague initializes a league object
func NewLeague(r io.Reader) ([]Player, error) {
	var league []Player
	err := json.NewDecoder(r).Decode(&league)
	if err != nil {
		err = fmt.Errorf("problem parsing league, %v", err)
	}

	return league, err
}

// Find retrieve a player if they exist in the league
// Returns nil if player doesn't exist
func (l League) Find(name string) *Player {
	for i, p := range l {
		if p.Name == name {
			return &l[i]
		}
	}
	return nil
}
