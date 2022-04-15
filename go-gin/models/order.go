package models

type Order struct {
	ID         OrderId
	Products   []ProductId
	TotalPrice Price
	CustomerId UserId
}
