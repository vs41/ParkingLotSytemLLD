package model

import "fmt"

type PaymentStatus string

const (
	PaymentPending   PaymentStatus = "PENDING"
	PaymentSucceeded PaymentStatus = "SUCCESS"
	PaymentFailed    PaymentStatus = "FAILED"
)

type Payment struct {
	PaymentID     string
	Amount        float64
	TicketID      string
	Method        string
	Status        PaymentStatus
	TransactionID string
}

func NewPayment(paymentID string, amount float64, ticketID, method string, status PaymentStatus, transactionID string) *Payment {
	return &Payment{PaymentID: paymentID, Amount: amount, TicketID: ticketID, Method: method, Status: status, TransactionID: transactionID}
}

func (p *Payment) GetPaymentID() string     { return p.PaymentID }
func (p *Payment) GetAmount() float64       { return p.Amount }
func (p *Payment) GetTicketID() string      { return p.TicketID }
func (p *Payment) GetMethod() string        { return p.Method }
func (p *Payment) GetStatus() PaymentStatus { return p.Status }
func (p *Payment) GetTransactionID() string { return p.TransactionID }

func (p *Payment) MarkSuccessful(transactionID string) error {
	if p == nil || p.Status != PaymentPending {
		return fmt.Errorf("only a pending payment can be marked successful")
	}
	if transactionID == "" {
		return fmt.Errorf("transaction ID is required")
	}
	p.TransactionID = transactionID
	p.Status = PaymentSucceeded
	return nil
}

func (p *Payment) MarkFailed() error {
	if p == nil || p.Status != PaymentPending {
		return fmt.Errorf("only a pending payment can be marked failed")
	}
	p.Status = PaymentFailed
	return nil
}
