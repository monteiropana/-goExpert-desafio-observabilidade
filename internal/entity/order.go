package entity

import "errors"

// regras de negocios sao essas
// precisa de uma camada para ter a intencao do usuario
type Order struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

func NewOrder(id string, price, Tax float64) (*Order, error) {
	order := &Order{
		ID:    id,
		Price: price,
		Tax:   Tax,
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
	if o.Tax <= 0 {
		return errors.New("invalid Tax")
	}
	return nil
}

func (o *Order) CalculaPrecoFInal() error {
	o.FinalPrice = o.Price + o.Tax
	err := o.IsValid()
	if err != nil {
		return err
	}
	return nil
}
