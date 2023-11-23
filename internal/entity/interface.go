package entity

type OrderRepositoryInterface interface {
	Save(order *Order) error
	GetTotalTranscations() (int, error)
}
