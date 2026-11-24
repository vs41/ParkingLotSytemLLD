package repository

import (
	"fmt"
	"parkinglot/database"
	"parkinglot/model"
)

type PaymentRepository struct{ database *database.Database }

func NewPaymentRepository(db *database.Database) *PaymentRepository {
	return &PaymentRepository{database: db}
}
func (p *PaymentRepository) SavePayment(payment model.Payment) error {
	if payment.PaymentID == "" {
		return fmt.Errorf("payment ID is required")
	}
	p.database.SavePayment(payment.PaymentID, payment)
	return nil
}
func (p *PaymentRepository) GetPaymentByID(paymentID string) (model.Payment, bool) {
	return p.database.GetPayment(paymentID)
}
