package service

import (
	"fmt"
	"strings"

	"parkinglot/model"
	"parkinglot/utilies"
)

// PaymentStrategy is implemented by every supported payment method.
type PaymentStrategy interface {
	Pay(amount float64, ticketID string) error
}

type PaymentRepository interface {
	SavePayment(payment model.Payment) error
}

type UPIPaymentService struct{ paymentRepo PaymentRepository }
type CardPaymentService struct{ paymentRepo PaymentRepository }
type CashPaymentService struct{ paymentRepo PaymentRepository }

func (u *UPIPaymentService) Pay(amount float64, ticketID string) error {
	return saveSuccessfulPayment(u.paymentRepo, amount, ticketID, "UPI")
}
func (c *CardPaymentService) Pay(amount float64, ticketID string) error {
	return saveSuccessfulPayment(c.paymentRepo, amount, ticketID, "CARD")
}
func (c *CashPaymentService) Pay(amount float64, ticketID string) error {
	return saveSuccessfulPayment(c.paymentRepo, amount, ticketID, "CASH")
}

func saveSuccessfulPayment(repo PaymentRepository, amount float64, ticketID, method string) error {
	if amount < 0 {
		return fmt.Errorf("payment amount cannot be negative")
	}
	if ticketID == "" {
		return fmt.Errorf("ticket ID is required")
	}
	payment := model.NewPayment(
		utilies.NewIdGenerator().GenerateID(), amount, ticketID, method,
		model.PaymentPending, "",
	)
	// This demo simulates a successful provider response. In a real integration,
	// call MarkSuccessful only after the payment gateway confirms the transaction.
	if err := payment.MarkSuccessful(utilies.NewIdGenerator().GenerateID()); err != nil {
		return err
	}
	return repo.SavePayment(*payment)
}

// PaymentService owns the strategy registry and chooses a strategy per request.
type PaymentService struct{ strategies map[string]PaymentStrategy }

func NewPaymentService(repo PaymentRepository) *PaymentService {
	return &PaymentService{strategies: map[string]PaymentStrategy{
		"UPI":  &UPIPaymentService{paymentRepo: repo},
		"CARD": &CardPaymentService{paymentRepo: repo},
		"CASH": &CashPaymentService{paymentRepo: repo},
	}}
}

func (p *PaymentService) ProcessPayment(method string, amount float64, ticketID string) error {
	if p == nil {
		return fmt.Errorf("payment service is not configured")
	}
	key := strings.ToUpper(strings.TrimSpace(method))
	strategy, ok := p.strategies[key]
	if !ok {
		return fmt.Errorf("unsupported payment method %q", method)
	}
	return strategy.Pay(amount, ticketID)
}
