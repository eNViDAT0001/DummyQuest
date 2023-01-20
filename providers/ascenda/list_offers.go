package ascenda

import (
	"DummyQuest/providers/ascenda/io"
	"DummyQuest/providers/base_api"
	"encoding/json"
	"fmt"
	"net/http"
)

func (p *ascendaProvider) ListOffers(latitude float32, longitude float32, radius float32) ([]io.AscendaOffer, error) {
	url := MockHost + "/offers/near_by"

	query := ""
	if latitude > 0 {
		query += fmt.Sprintf("lat=%f&", latitude)
	}
	if longitude > 0 {
		query += fmt.Sprintf("lon=%f&", longitude)
	}
	if radius > 0 {
		query += fmt.Sprintf("rad=%f", radius)
	}

	api := &base_api.BaseAPI{
		Header:  nil,
		Url:     url,
		Query:   query,
		Payload: nil,
		Client:  http.DefaultClient,
	}

	resp, err := api.Request(http.MethodGet)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var result io.AscendaOfferRes
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.Offers, nil
}
