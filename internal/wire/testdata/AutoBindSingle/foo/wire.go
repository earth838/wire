//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
)

func injectedMessage() I {
	wire.Build(
		provideS1)
	return nil
}
