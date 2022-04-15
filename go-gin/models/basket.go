package models

type Basket struct {
	ID       BasketId
	Products []ProductId
}

func (bsk Basket) AddProd(pId ProductId) {
	// smth
}

func (bsk Basket) RemoveProd(pId ProductId) {
	// smth
}
