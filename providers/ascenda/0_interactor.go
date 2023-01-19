package ascenda

import (
	"DummyQuest/providers/ascenda/io"
)

const (
	MockHost = "https://61c3deadf1af4a0017d990e7.mockapi.io"
)

type AscendaProvider interface {
	ListOffers(latitude float32, longitude float32, radius float32) ([]io.AscendaOffer, error)
}
type ascendaProvider struct {
}

func NewAscendaProvider() AscendaProvider {
	return &ascendaProvider{}
}
