package app

import "context"

type App struct{}

func New(ctx context.Context) (*App, error) {
	a := &App{}

	return a, nil
}
