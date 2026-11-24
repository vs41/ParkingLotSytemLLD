package database

import (
	"parkinglot/model"
	"sync"
)

type Database struct {
	mu       sync.RWMutex
	Tickets  map[string]model.Ticket
	Payments map[string]model.Payment
}

func NewDatabase() *Database {
	return &Database{Tickets: make(map[string]model.Ticket), Payments: make(map[string]model.Payment)}
}

func (d *Database) SaveTicket(id string, ticket model.Ticket) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.Tickets[id] = ticket
}
func (d *Database) GetTicket(id string) (model.Ticket, bool) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	ticket, ok := d.Tickets[id]
	return ticket, ok
}
func (d *Database) UpdateTicket(id string, ticket model.Ticket) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.Tickets[id] = ticket
}
func (d *Database) DeleteTicket(id string) {
	d.mu.Lock()
	defer d.mu.Unlock()
	delete(d.Tickets, id)
}
func (d *Database) SavePayment(id string, payment model.Payment) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.Payments[id] = payment
}
func (d *Database) GetPayment(id string) (model.Payment, bool) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	payment, ok := d.Payments[id]
	return payment, ok
}
func (d *Database) DeletePayment(id string) {
	d.mu.Lock()
	defer d.mu.Unlock()
	delete(d.Payments, id)
}
