package data

import (
	"time"
)

type Product struct {
	ID int
	Name string
	Description string
	Price float32
	SKU string
	CreatedOn string
	UpdatedOn string
	DeletedOn string
}