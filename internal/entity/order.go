package entity

import "errors"

type Order struct {
	ID         string
	Price      float64
	Juros      float64
	FinalPrice float64
}

func NewOrder(id string, price, juros float64) (*Order, error) {
	order := &Order{
		ID:    id,
		Price: price,
		Juros: juros,
	}
	err := order.IsValid()
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (o *Order) IsValid() error {
	if o.ID == "" {
		return errors.New("ID is required")
	}
	if o.Price <= 0 {
		return errors.New("invalid price")
	}
	if o.Juros <= 0 {
		return errors.New("invalid juros")
	}
	return nil
}

func (o *Order) CalculaPrecoFInal() error {
	o.FinalPrice = o.Price + o.Juros
	err := o.IsValid()
	if err != nil {
		return err
	}
	return nil
}
