package usecase

import (
	"github.com/desafio/clean-arch/internal/entity"
	EventsPkg "github.com/desafio/clean-arch/pkg/events"
)

// recebe os dados e orquestra o processo
// cria uma orden > insre no banco > e dispara o evento
type InputOderDTO struct {
	ID    string  `json:"id"`
	Price float64 `json:"price"`
	Juros float64 `json:"juros"`
}

type OutputOderDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Juros      float64 `json:"juros"`
	PriceFinal float64 `json:"price_final"`
}

type CreateOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
	OrderCreated    EventsPkg.EventInterface
	EventDispatcher EventsPkg.EventDispatcherInterface
}

func NewCreateOrderUseCase(
	OrderRepository entity.OrderRepositoryInterface,
	OrderCreated EventsPkg.EventInterface,
	EventDispatcher EventsPkg.EventDispatcherInterface,
) *CreateOrderUseCase {
	return &CreateOrderUseCase{
		OrderRepository: OrderRepository,
		OrderCreated:    OrderCreated,
		EventDispatcher: EventDispatcher,
	}
}

func (c *CreateOrderUseCase) Execute(input InputOderDTO) (OutputOderDTO, error) {
	order := entity.Order{
		ID:    input.ID,
		Price: input.Price,
		Juros: input.Juros,
	}
	order.CalculaPrecoFInal()
	if err := c.OrderRepository.SaveOrder(&order); err != nil {
		return OutputOderDTO{}, err
	}

	dto := OutputOderDTO{
		ID:         order.ID,
		Price:      order.Price,
		Juros:      order.Juros,
		PriceFinal: order.FinalPrice,
	}

	c.OrderCreated.SetPayload(dto)
	c.EventDispatcher.Dispatch(c.OrderCreated)

	return dto, nil
}
