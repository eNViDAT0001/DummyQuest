package io

import "DummyQuest/providers/ascenda/io"

type ListOffersByUserCheckinRes struct {
	Offers []io.AscendaOffer `json:"offers"`
}
