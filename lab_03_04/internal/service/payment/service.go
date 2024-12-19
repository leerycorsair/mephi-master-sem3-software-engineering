package payment

//go:generate mockgen -destination mocks/storage.go -package mocks . Storage
//go:generate mockgen -destination mocks/menu.go -package mocks . Menu

import (
	"context"
	"fmt"
	"sync"

	"github.com/go-faster/errors"
)

type Menu interface {
	GetItemPrice(ctx context.Context, itemID int) (float64, error)
}

type Storage interface {
	CreatePayment(ctx context.Context, orderID int, amount float64, method string) error
	GetPayment(ctx context.Context, orderID int) (Payment, error)
	UpdatePayment(ctx context.Context, orderID int, status string) error
	DeletePayment(ctx context.Context, orderID int) error
}

type Service struct {
	storage Storage
	menu    Menu
}

func New(storage Storage, menu Menu) *Service {
	return &Service{
		storage: storage,
		menu:    menu,
	}
}

func (s *Service) CalculateTotal(ctx context.Context, orderID int, items []OrderItem) (float64, error) {
	const op = "service.CalculateTotal"

	total := 0.0
	var wg sync.WaitGroup
	errChan := make(chan error, len(items))
	priceChan := make(chan float64, len(items))

	for _, item := range items {
		wg.Add(1)
		go func(item OrderItem) {
			defer wg.Done()
			price, err := s.menu.GetItemPrice(ctx, item.MenuItemID)
			if err != nil {
				errChan <- fmt.Errorf("fetch price for item %d: %v", item.MenuItemID, err)
				return
			}
			priceChan <- float64(item.Quantity) * price
		}(item)
	}

	wg.Wait()
	close(errChan)
	close(priceChan)

	for err := range errChan {
		return 0, errors.Wrap(err, op)
	}
	for price := range priceChan {
		total += price
	}

	return total, nil
}

func (s *Service) ProcessPayment(ctx context.Context, orderID int, items []OrderItem, method string) (bool, error) {
	const op = "service.ProcessPayment"

	total, err := s.CalculateTotal(ctx, orderID, items)
	if err != nil {
		return false, errors.Wrap(err, op)
	}

	err = s.storage.CreatePayment(ctx, orderID, total, method)
	if err != nil {
		return false, errors.Wrap(err, op)
	}

	return true, nil
}

func (s *Service) RefundPayment(ctx context.Context, orderID int) error {
	const op = "service.RefundPayment"

	payment, err := s.storage.GetPayment(ctx, orderID)
	if err != nil {
		return errors.Wrap(err, op)
	}

	if payment.Status != PaymentCompleted {
		return fmt.Errorf("%s: status is %s", op, payment.Status)
	}

	err = s.storage.UpdatePayment(ctx, orderID, PaymentRefunded)
	if err != nil {
		return errors.Wrap(err, op)
	}

	return nil
}
