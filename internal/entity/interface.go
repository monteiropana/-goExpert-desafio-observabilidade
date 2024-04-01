// Entidades regras de negocios
package entity

//Regras de negocios
//coracao da application

type RepositoryOrderInterface interface {
	SaveOrder(order *Order) error
	// GetTotal() (int, error)
	GetListOfOrders() ([]*Order, error)
}
