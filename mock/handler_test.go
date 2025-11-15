package main

import (
	"context"
	"testing"
)

type UserTeamGetterMock struct {
	getUserFunc func(id int) (*User, error)
	getTeamFunc func(id int) (*Team, error)
}

func (m *UserTeamGetterMock) GetUser(id int) (*User, error) {
	if m != nil && m.getUserFunc != nil {
		return m.getUserFunc(id)
	}
	// Otherwise return a default value or error
	return &User{ID: id, Name: "DefaultUser", DefaultTeamID: 999}, nil
}

func (m *UserTeamGetterMock) GetTeam(id int) (*Team, error) {
	if m != nil && m.getTeamFunc != nil {
		return m.getTeamFunc(id)
	}
	// Otherwise return a default value or error
	return &Team{ID: id, Name: "DefaultTeam"}, nil
}

func TestSomeHandler(t *testing.T) {

	tests := []struct {
		name    string
		store   *UserTeamGetterMock
		wantErr bool
	}{
		{
			name:    "happy path (nil mock uses defaults)",
			store:   nil,
			wantErr: false,
		},
		{
			name: "GetUser returns error",
			store: &UserTeamGetterMock{
				getUserFunc: func(id int) (*User, error) {
					return nil, ErrUserNotFound
				},
			},
			wantErr: true,
		},
		{
			name: "GetTeam returns error - 2nd call",
			store: &UserTeamGetterMock{
				getTeamFunc: func(id int) (*Team, error) {
					return nil, ErrTeamNotFound
				},
			},
			wantErr: true,
		},
	}

	ctx := context.Background()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := SomeHandler(ctx, tt.store, 1)
			if (err != nil) != tt.wantErr {
				t.Errorf("SomeHandler() error = %v, wantErr %v", err, tt.wantErr)
			}

		})

	}

}
