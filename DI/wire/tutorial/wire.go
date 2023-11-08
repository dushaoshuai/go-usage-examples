//go:build wireinject

package main

import "github.com/google/wire"

func initializeEvent(phrase string) (Event, error) {
	wire.Build(NewEvent, NewGreeter, NewMessage)
	return Event{}, nil
}
