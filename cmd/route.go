package main

import "github.com/gin-gonic/gin"

func router(r *gin.Engine) {
	allHandler := initHandlerCollection()
	v1 := r.Group("/api/v1")
	{
		offerGroup := v1.Group("/offers/near_by")
		{
			offerGroup.GET("", allHandler.offerHttpHandler.ListOffers())
		}
	}
}
