package storage

type OfferStorage interface {
}

type offerStorage struct {
}

func NewOfferStorage() OfferStorage {
	return &offerStorage{}
}
