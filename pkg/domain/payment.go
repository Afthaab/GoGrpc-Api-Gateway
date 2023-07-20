package domain

import "github.com/google/uuid"

type PaymentData struct {
	ID          uint    `json:"id"`
	Totalamount float32 `JSON:"totalamount"`
	Couponid    string  `JSON:"couponid"`
}

type Orders struct {
	Paymentid int64 `json:"paymentid"`
	Addid     int64 `json:"addid"`
}

type OrderItems struct {
	Orderid     uuid.UUID     `json:"orderid"`
	Cartproduct []CartProduct `json:"cartproduct"`
}
