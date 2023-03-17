package main

import (
	"errors"
	"fmt"
)

type Order struct {
	WarehouseFilled bool
	ShippingFilled  bool
	BillingFilled   bool
}

type OrderHandler interface {
	SetNext(OrderHandler) OrderHandler
	Handle(*Order) error
}

type OrderProcessor struct {
	handler OrderHandler
}

func (o *OrderProcessor) SetHandler(handler OrderHandler) {
	o.handler = handler
}

func (o *OrderProcessor) Process(order *Order) error {
	return o.handler.Handle(order)
}

type WarehouseHandler struct {
	next OrderHandler
}

func (w *WarehouseHandler) SetNext(next OrderHandler) OrderHandler {
	w.next = next
	return next
}

func (w *WarehouseHandler) Handle(order *Order) error {
	if order.WarehouseFilled {
		if w.next != nil {
			return w.next.Handle(order)
		}
		return nil
	}
	return errors.New("warehouse not filled")
}

type ShippingHandler struct {
	next OrderHandler
}

func (s *ShippingHandler) SetNext(next OrderHandler) OrderHandler {
	s.next = next
	return next
}

func (s *ShippingHandler) Handle(order *Order) error {
	if order.ShippingFilled {
		if s.next != nil {
			return s.next.Handle(order)
		}
		return nil
	}
	return errors.New("shipping not filled")
}

type BillingHandler struct {
	next OrderHandler
}

func (b *BillingHandler) SetNext(next OrderHandler) OrderHandler {
	b.next = next
	return next
}

func (b *BillingHandler) Handle(order *Order) error {
	if order.BillingFilled {
		if b.next != nil {
			return b.next.Handle(order)
		}
		return nil
	}
	return errors.New("billing not filled")
}

func main() {
	order := &Order{
		WarehouseFilled: true,
		ShippingFilled:  true,
		BillingFilled:   true,
	}

	orderProcessor := &OrderProcessor{}
	warehouseHandler := &WarehouseHandler{}
	shippingHandler := &ShippingHandler{}
	billingHandler := &BillingHandler{}

	warehouseHandler.SetNext(shippingHandler)
	shippingHandler.SetNext(billingHandler)

	orderProcessor.SetHandler(warehouseHandler)

	if err := orderProcessor.Process(order); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Order processed successfully")
}
