package payment_test

import (
	"context"
	"testing"

	"github.com/go-faster/errors"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"

	"restoraunt/internal/service/payment"
	"restoraunt/internal/service/payment/mocks"
)

type ServiceSuite struct {
	suite.Suite
	service *payment.Service
	storage *mocks.MockStorage
	menu    *mocks.MockMenu
}

func (s *ServiceSuite) SetupTest() {
	ctrl := gomock.NewController(s.T())
	s.storage = mocks.NewMockStorage(ctrl)
	s.menu = mocks.NewMockMenu(ctrl)
	s.service = payment.New(
		s.storage,
		s.menu,
	)
}

func (s *ServiceSuite) TestCalculateTotal_Success() {
	ctx := context.Background()
	items := []payment.OrderItem{
		{MenuItemID: 1, Quantity: 2},
		{MenuItemID: 2, Quantity: 3},
	}

	s.menu.EXPECT().GetItemPrice(ctx, 1).Return(10.0, nil)
	s.menu.EXPECT().GetItemPrice(ctx, 2).Return(20.0, nil)

	total, err := s.service.CalculateTotal(ctx, 1, items)
	s.NoError(err)
	s.Equal(80.0, total)
}

func (s *ServiceSuite) TestCalculateTotal_ErrorFetchingPrice() {
	ctx := context.Background()
	items := []payment.OrderItem{
		{MenuItemID: 1, Quantity: 2},
	}

	s.menu.EXPECT().GetItemPrice(ctx, 1).Return(0.0, context.DeadlineExceeded)

	total, err := s.service.CalculateTotal(ctx, 1, items)
	s.EqualError(err, "service.CalculateTotal: fetch price for item 1: context deadline exceeded")
	s.Zero(total)
}

func (s *ServiceSuite) TestProcessPayment_Success() {
	ctx := context.Background()
	items := []payment.OrderItem{
		{MenuItemID: 1, Quantity: 2},
	}

	s.menu.EXPECT().GetItemPrice(ctx, 1).Return(10.0, nil)
	s.storage.EXPECT().CreatePayment(ctx, 1, 20.0, "card").Return(nil)

	success, err := s.service.ProcessPayment(ctx, 1, items, "card")
	s.NoError(err)
	s.True(success)
}

func (s *ServiceSuite) TestProcessPayment_ErrorCreatingPayment() {
	ctx := context.Background()
	items := []payment.OrderItem{
		{MenuItemID: 1, Quantity: 2},
	}

	s.menu.EXPECT().GetItemPrice(ctx, 1).Return(10.0, nil)
	s.storage.EXPECT().CreatePayment(ctx, 1, 20.0, "card").Return(context.DeadlineExceeded)

	success, err := s.service.ProcessPayment(ctx, 1, items, "card")
	s.EqualError(err, "service.ProcessPayment: context deadline exceeded")
	s.Zero(success)
}

func (s *ServiceSuite) TestProcessPayment_ErrorCalculateTotal() {
	ctx := context.Background()
	items := []payment.OrderItem{
		{MenuItemID: 1, Quantity: 2},
	}

	s.menu.EXPECT().GetItemPrice(ctx, 1).Return(10.0, errors.New("some error"))

	success, err := s.service.ProcessPayment(ctx, 1, items, "card")
	s.EqualError(err, "service.ProcessPayment: service.CalculateTotal: fetch price for item 1: some error")
	s.False(success)
}

func (s *ServiceSuite) TestRefundPayment_Success() {
	ctx := context.Background()
	paymentRecord := payment.Payment{
		OrderID: 1,
		Amount:  20.0,
		Method:  "card",
		Status:  payment.PaymentCompleted,
	}

	s.storage.EXPECT().GetPayment(ctx, 1).Return(paymentRecord, nil)
	s.storage.EXPECT().UpdatePayment(ctx, 1, payment.PaymentRefunded).Return(nil)

	err := s.service.RefundPayment(ctx, 1)
	s.NoError(err)
}

func (s *ServiceSuite) TestRefundPayment_ErrorInvalidStatus() {
	ctx := context.Background()
	paymentRecord := payment.Payment{
		OrderID: 1,
		Amount:  20.0,
		Method:  "card",
		Status:  payment.PaymentPending,
	}

	s.storage.EXPECT().GetPayment(ctx, 1).Return(paymentRecord, nil)

	err := s.service.RefundPayment(ctx, 1)
	s.EqualError(err, "service.RefundPayment: status is Pending")
}

func (s *ServiceSuite) TestRefundPayment_ErrorUpdatePayment() {
	ctx := context.Background()
	paymentRecord := payment.Payment{
		OrderID: 1,
		Amount:  20.0,
		Method:  "card",
		Status:  payment.PaymentCompleted,
	}

	s.storage.EXPECT().GetPayment(ctx, 1).Return(paymentRecord, nil)
	s.storage.EXPECT().UpdatePayment(ctx, 1, payment.PaymentRefunded).Return(errors.New("some error"))

	err := s.service.RefundPayment(ctx, 1)
	s.EqualError(err, "service.RefundPayment: some error")
}

func (s *ServiceSuite) TestRefundPayment_ErrorGetPayment() {
	ctx := context.Background()

	s.storage.EXPECT().GetPayment(ctx, 1).Return(payment.Payment{}, errors.New("some error"))

	err := s.service.RefundPayment(ctx, 1)
	s.EqualError(err, "service.RefundPayment: some error")
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}
