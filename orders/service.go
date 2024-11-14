package main

import "context"

// สร้าง Struct ชื่อ service ที่มี field store ที่เป็น OrderStore interface จาก types.go
type service struct {
	store OrderStore
}

// สร้าง function ชื่อ newService ที่รับ OrderStore และ return pointer ของ service
func newService(store OrderStore) *service {
	// สร้าง instance ของ service และ return ออกไป
	return &service{store}
}

func (s *service) CreateOrder(context.Context) error {
	return nil
}