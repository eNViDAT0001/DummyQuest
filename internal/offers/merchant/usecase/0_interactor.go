package usecase

import (
	"DummyQuest/internal/offers/merchant/storage"
	"DummyQuest/providers/ascenda"
	"DummyQuest/providers/ascenda/io"
	"context"
	"time"
)

type OfferUseCase interface {
	ListOffers(ctx *context.Context, latitude float32, longitude float32, radius float32, checkinDate *time.Time) ([]io.AscendaOffer, error)
}

type offerUseCase struct {
	offerStorage storage.OfferStorage
	ascendaAPI   ascenda.AscendaProvider
}

func NewOfferUseCase(
	offerStorage storage.OfferStorage,
	ascendaAPI ascenda.AscendaProvider) OfferUseCase {
	return &offerUseCase{
		offerStorage: offerStorage,
		ascendaAPI:   ascendaAPI,
	}
}
