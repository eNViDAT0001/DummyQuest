// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"DummyQuest/delivery/offer"
	"DummyQuest/internal/offers/merchant/storage"
	"DummyQuest/internal/offers/merchant/usecase"
	"DummyQuest/providers/ascenda"
)

// Injectors from wire.go:

func initHandlerCollection() *HandlerCollection {
	offerStorage := storage.NewOfferStorage()
	ascendaProvider := ascenda.NewAscendaProvider()
	offerUseCase := usecase.NewOfferUseCase(offerStorage, ascendaProvider)
	offerHttpHandler := handler.NewOfferHttpHandler(offerUseCase)
	handlerCollection := NewHandlerCollection(offerHttpHandler)
	return handlerCollection
}