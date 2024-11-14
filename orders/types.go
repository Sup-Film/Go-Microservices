package main

import "context"

// สร้าง interface สำหรับ OrderService ที่มี method CreateOrder ที่รับ context.Context และ return error
type OrderService interface {
	CreateOrder(context.Context) error
}

// สร้าง interface สำหรับ OrderStore ที่มี method Create ที่รับ context.Context และ return error
type OrderStore interface {
	Create(context.Context) error
}
