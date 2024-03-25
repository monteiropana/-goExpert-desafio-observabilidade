// Entidades regras de negocios
package entity

type RepositoryOrder interface {
	Save(order *Order) error
	//GetListOrders
	//GetTotalOrders
}
