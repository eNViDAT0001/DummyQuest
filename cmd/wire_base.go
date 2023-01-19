package main

import (
	offerHttpHandlerPKG "DummyQuest/delivery/offer"
	offerStoPKG "DummyQuest/internal/offers/merchant/storage"
	offerUCPKG "DummyQuest/internal/offers/merchant/usecase"

	ascendaPKG "DummyQuest/providers/ascenda"

	"github.com/google/wire"
)

var IteratorCollection = wire.NewSet(
	offerHttpHandlerPKG.NewOfferHttpHandler,
	offerUCPKG.NewOfferUseCase,
	offerStoPKG.NewOfferStorage,

	ascendaPKG.NewAscendaProvider,

	NewHandlerCollection,
)

type HandlerCollection struct {
	offerHttpHandler offerHttpHandlerPKG.OfferHttpHandler
}

func NewHandlerCollection(
	offerHttpHandler offerHttpHandlerPKG.OfferHttpHandler,
) *HandlerCollection {
	return &HandlerCollection{
		offerHttpHandler: offerHttpHandler,
	}
}
