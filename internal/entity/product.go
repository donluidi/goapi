package entity

import (
	"goapi/pkg"
	"time"
)

type product struct {
	id        pkg.Id    `json:"id"`
	name      string    `json:"name"`
	price     int       `json:"price"`
	createdAt time.Time `json:"created_at"`
}

func NewProduct(new_name string, new_price int) *product {
	return &product{
		id:        pkg.NewId(),
		name:      new_name,
		price:     new_price,
		createdAt: time.Now(),
	}
}

func (p *product) GetId() pkg.Id {
	return p.id
}

func (p *product) GetName() string {
	return p.name
}

func (p *product) GetPrice() int {
	return p.price
}

func (p *product) GetCreatedAt() time.Time {
	return p.createdAt
}
