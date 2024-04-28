package usecase

import (
	"github.com/desafio/clean-arch/internal/entity"
	"github.com/desafio/clean-arch/pkg/events"
)

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
	OrdersListed    events.EventInterface
	EventDispatcher events.EventDispatcherInterface
}

func NewListOrdersUseCase(OrderRepository entity.OrderRepositoryInterface, OrdersListed events.EventInterface, EventDispatcher events.EventDispatcherInterface) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: OrderRepository,
		OrdersListed:    OrdersListed,
		EventDispatcher: EventDispatcher,
	}
}

func (l *ListOrdersUseCase) Execute() ([]OutputOderDTO, error) {
	orders, err := l.OrderRepository.GetListOfOrders()
	if err != nil {
		return nil, err
	}

	var dtos []OutputOderDTO
	for _, order := range orders {
		dto := OutputOderDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		}
		dtos = append(dtos, dto)
	}

	payload := map[string]interface{}{
		"orders": dtos,
	}

	l.OrdersListed.SetPayload(payload)
	l.EventDispatcher.Dispatch(l.OrdersListed)

	return dtos, nil
}
