package io

import (
	"DummyQuest/external/common"
)

type AscendaOfferReq struct {
	Offers []AscendaOffer `json:"offers"`
}
type AscendaOffer struct {
	ID        int             `json:"id"`
	Title     string          `json:"title"`
	Category  int             `json:"category"`
	Merchants []Merchant      `json:"merchants"`
	ValidTo   common.DateOnly `json:"valid_to"`
}

type Merchant struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Distance float64 `json:"distance"`
}
