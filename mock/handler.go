package main

import "context"

type UserGetter interface {
	GetUser(id int) (*User, error)
}

type TeamGetter interface {
	GetTeam(id int) (*Team, error)
}

type UserTeamGetter interface {
	UserGetter
	TeamGetter
}

func SomeHandler(ctx context.Context, store UserTeamGetter, UserId int) error {
	user, err := store.GetUser(UserId)
	if err != nil {
		return err
	}
	_, err = store.GetTeam(user.DefaultTeamID)
	if err != nil {
		return err
	}
	return nil
}
