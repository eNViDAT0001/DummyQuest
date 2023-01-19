package usecase

import (
	"DummyQuest/providers/ascenda/io"
	"context"
	"time"
)

const (
	MinimumDate   = 5
	HotelCategory = 3
)

func (uc *offerUseCase) ListOffers(ctx *context.Context, latitude float32, longitude float32, radius float32, checkinDate *time.Time) ([]io.AscendaOffer, error) {
	ascendaOffers, err := uc.ascendaAPI.ListOffers(latitude, longitude, radius)
	if err != nil {
		return nil, err
	}

	validOffers := make([]io.AscendaOffer, 0)
	validDate := checkinDate.AddDate(0, 0, MinimumDate)
	for idx, offer := range ascendaOffers {
		// 1 - Only select category Restaurant , Retail & Activity offers.
		isHotelOffer := offer.Category == HotelCategory
		// 2 - Offer need to be valid till checkin date + 5 days. (valid_to is in YYYY-MM-DD)
		if offer.ValidTo.Before(validDate) || isHotelOffer {
			continue
		}

		// 3 - If an offer is available in multiple merchants, only select the closest merchant
		if len(offer.Merchants) > 1 {
			//Get nearest Merchant Index
			markerIndex := 0
			for i, merchant := range offer.Merchants {
				if merchant.Distance < offer.Merchants[markerIndex].Distance {
					markerIndex = i
				}
			}
			ascendaOffers[idx].Merchants = []io.Merchant{offer.Merchants[markerIndex]}
		}

		// 5 - Both final selected offers should be in different categories.
		//     If there are multiple offers in the same category give priority to the closest merchant offer.
		categoryDuplicated := false
		for i, validOffer := range validOffers {
			if validOffer.Category == offer.Category {
				categoryDuplicated = true

				if validOffer.Merchants[0].Distance > offer.Merchants[0].Distance {
					validOffers[i] = offer
				}

				break
			}
		}

		if categoryDuplicated {
			continue
		}

		validOffers = append(validOffers, ascendaOffers[idx])
	}

	// 4 - This class should only return 2 offers even though there are several eligible offers
	// 6 - If there are multiple offers with different categories,
	//     select the closest merchant offers when selecting 2 offers
	if len(validOffers) > 2 {
		// Sort first, get validOffers[0] and validOffers[1] as 2 nearest offers
		//TODO: Update sort algorithm
		//     or replace with find Min(st) and Min(nd) only
		//     when we have huge categories
		for i, _ := range validOffers {
			minimumIdx := i
			for j := i; j < len(validOffers); j++ {
				if validOffers[minimumIdx].Merchants[0].Distance > validOffers[j].Merchants[0].Distance {
					minimumIdx = j
				}
			}
			validOffers[i], validOffers[minimumIdx] = validOffers[minimumIdx], validOffers[i]
		}

	}

	return validOffers[0:2], nil
}
