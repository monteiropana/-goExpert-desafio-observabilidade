package usecase

type InputOderDTO struct {
	ID       string  `json:"id"`
	Price    float64 `json:"price"`
	Juros    float64 `json:"juros"`
	Quantity int     `json:"quantity"`
}

type OutputOderDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Juros      float64 `json:"juros"`
	Quantity   int     `json:"quantity"`
	PriceFinal float64 `json:"price_final"`
}
