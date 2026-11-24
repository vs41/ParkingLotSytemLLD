package repository

import (
	"fmt"
	"parkinglot/database"
	"parkinglot/model"
)

type TicketRepository struct{ database *database.Database }

func NewTicketRepository(db *database.Database) *TicketRepository {
	return &TicketRepository{database: db}
}
func (t *TicketRepository) SaveTicket(ticket model.Ticket) error {
	if ticket.TicketID == "" {
		return fmt.Errorf("ticket ID is required")
	}
	t.database.SaveTicket(ticket.TicketID, ticket)
	return nil
}
func (t *TicketRepository) GetTicketByID(id string) (model.Ticket, bool) {
	return t.database.GetTicket(id)
}
func (t *TicketRepository) UpdateTicket(id string, ticket model.Ticket) error {
	if _, ok := t.database.GetTicket(id); !ok {
		return fmt.Errorf("ticket %q not found", id)
	}
	t.database.UpdateTicket(id, ticket)
	return nil
}
func (t *TicketRepository) DeleteTicket(id string) { t.database.DeleteTicket(id) }
