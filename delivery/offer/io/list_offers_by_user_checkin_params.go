package io

type ListOffersParams struct {
	Latitude  float32 `form:"lat"`
	Longitude float32 `form:"lon"`
	Radius    float32 `form:"rad"`
	Checkin   string  `form:"checkin" binding:"required"`
}
