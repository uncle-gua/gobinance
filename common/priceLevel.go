package common

import (
	"encoding/json"
	"errors"
	"strconv"
)

// PriceLevel is a common structure for bids and asks in the
// order book.

var ErrPriceLevel = errors.New("failed to parse price level")

type PriceLevel struct {
	Price    float64
	Quantity float64
}

func (p *PriceLevel) UnmarshalJSON(data []byte) error {
	var items []string
	if err := json.Unmarshal(data, &items); err != nil {
		return err
	}

	if len(items) != 2 {
		return ErrPriceLevel
	}

	price, err := strconv.ParseFloat(items[0], 64)
	if err != nil {
		return err
	}

	Quantity, err := strconv.ParseFloat(items[1], 64)
	if err != nil {
		return err
	}

	p.Price = price
	p.Quantity = Quantity

	return nil
}

func (p *PriceLevel) MarshalJSON() ([]byte, error) {
	items := [2]string{}
	items[0] = strconv.FormatFloat(p.Price, 'f', -1, 64)
	items[1] = strconv.FormatFloat(p.Quantity, 'f', -1, 64)

	return json.Marshal(items)
}
