package entity

import (
	"time"
)

type Customers struct {
	Id        int
	Name      string
	Phone     string
	EntryDate time.Time
	OutDate   time.Time
	Bill      float64
}

type Services struct {
	Id    int
	Name  string
	Price float64
}

type Transactions struct {
	TransactionId int
	CustomerId    int
	ServiceId     int
	Quantity      int
	Unit          string
	DateEntry     time.Time
	Price         float64
	TotalPrice    float64
}
