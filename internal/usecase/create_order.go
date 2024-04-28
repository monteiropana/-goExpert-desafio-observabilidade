package usecase

import (
	"github.com/desafio/clean-arch/internal/entity"
	EventsPkg "github.com/desafio/clean-arch/pkg/events"
)

// recebe os dados e orquestra o processo
// cria uma orden > insre no banco > e dispara o evento
type OrderInputDTO struct {
	ID    string  `json:"id"`
	Price float64 `json:"price"`
	Tax   float64 `json:"Tax"`
}

type OutputOderDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"Tax"`
	FinalPrice float64 `json:"price_final"`
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

func (c *CreateOrderUseCase) Execute(input OrderInputDTO) (OutputOderDTO, error) {
	order := entity.Order{
		ID:    input.ID,
		Price: input.Price,
		Tax:   input.Tax,
	}
	order.CalculaPrecoFInal()
	if err := c.OrderRepository.Save(&order); err != nil {
		return OutputOderDTO{}, err
	}

	dto := OutputOderDTO{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
	}

	c.OrderCreated.SetPayload(dto)
	c.EventDispatcher.Dispatch(c.OrderCreated)

	return dto, nil
}
