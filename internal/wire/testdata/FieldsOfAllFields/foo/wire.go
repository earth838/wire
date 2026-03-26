//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
)

func injectedResult() Result {
	wire.Build(
		provideS,
		wire.FieldsOf(new(S), "*"),
		provideResult)
	return Result{}
}
