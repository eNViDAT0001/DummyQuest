//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"

func initHandlerCollection() *HandlerCollection {

	wire.Build(IteratorCollection)

	return &HandlerCollection{}
}
