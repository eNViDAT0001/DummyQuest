package handler

import (
	"DummyQuest/delivery/offer/io"
	"DummyQuest/external/common"
	"DummyQuest/external/request"
	"context"
	"github.com/gin-gonic/gin"
)

func (h *offerHttpHandler) ListOffers() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		var params io.ListOffersParams
		if err := cc.BindQuery(&params); err != nil {
			cc.ResponseError(err)
		}

		//convert string to date
		checkinDate, err := common.StringToDate(params.Checkin)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		offers, err := h.offerUseCase.ListOffers(&newCtx, params.Latitude, params.Longitude, params.Radius, checkinDate)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		cc.Ok(offers)
	}

}
