//go:build wireinject
// +build wireinject

package foo

import (
	"github.com/google/wire"
)

func injectedMessage() I {
	wire.Build(
		provideS1,
		provideS2)
	return nil
}
