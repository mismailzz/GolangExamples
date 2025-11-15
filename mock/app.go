package main

import (
	"context"
	"errors"
	"fmt"
)

var ErrUserNotFound = errors.New("user not found")
var ErrTeamNotFound = errors.New("team not found")

type User struct {
	ID            int
	Name          string
	DefaultTeamID int
}

type Team struct {
	ID   int
	Name string
}

type Store struct {
	usersClient map[int]User
	teamsClient map[int]Team
}

func NewStore() *Store {
	return &Store{
		usersClient: map[int]User{
			1: {ID: 1, Name: "Alice", DefaultTeamID: 101},
			2: {ID: 2, Name: "Bob", DefaultTeamID: 102},
		},
		teamsClient: map[int]Team{
			101: {ID: 101, Name: "Team Alpha"},
			102: {ID: 102, Name: "Team Beta"},
		},
	}
}

func (s *Store) GetUser(id int) (*User, error) {
	resp, exists := s.usersClient[id]
	if !exists {
		return nil, fmt.Errorf("For user id %d: %w", id, ErrUserNotFound)
	}
	return &resp, nil
}

func (s *Store) GetTeam(id int) (*Team, error) {
	resp, exists := s.teamsClient[id]
	if !exists {
		return nil, fmt.Errorf("For team id %d: %w", id, ErrTeamNotFound)
	}
	return &resp, nil
}

func main() {
	// Application entry point
	newStore := NewStore()

	u1, err := newStore.GetUser(1)
	fmt.Printf("User 1: %+v, err=%v\n", u1, err)

	u3, err := newStore.GetUser(3)
	fmt.Printf("User 3: %+v, err=%v\n", u3, err)

	t101, err := newStore.GetTeam(101)
	fmt.Printf("Team 101: %+v, err=%v\n", t101, err)

	t103, err := newStore.GetTeam(103)
	fmt.Printf("Team 103: %+v, err=%v\n", t103, err)

	ctx := context.Background()
	err = SomeHandler(ctx, newStore, 1)
	fmt.Printf("SomeHandler for User 1: err=%v\n", err)
}
