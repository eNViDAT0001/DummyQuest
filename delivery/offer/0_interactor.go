package handler

import (
	"DummyQuest/internal/offers/merchant/usecase"
	"github.com/gin-gonic/gin"
)

type OfferHttpHandler interface {
	ListOffers() func(ctx *gin.Context)
}

type offerHttpHandler struct {
	offerUseCase usecase.OfferUseCase
}

func NewOfferHttpHandler(offerUseCase usecase.OfferUseCase) OfferHttpHandler {
	return &offerHttpHandler{
		offerUseCase: offerUseCase,
	}
}
