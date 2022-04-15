package models

import "github.com/google/uuid"

type Price float32

type BasketId uuid.UUID

type OrderId uuid.UUID

type ProductId uuid.UUID

type UserId uuid.UUID

var strgProducts []Product = make([]Product, 0)
