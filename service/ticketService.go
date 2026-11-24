package service

import (
	"fmt"
	"parkinglot/model"
	"parkinglot/utilies"
	"time"
)

type TicketRepository interface {
	SaveTicket(model.Ticket) error
	GetTicketByID(string) (model.Ticket, bool)
	UpdateTicket(string, model.Ticket) error
	DeleteTicket(string)
}
type TicketService struct{ TicketRepository TicketRepository }

func NewTicketService(repo TicketRepository) *TicketService {
	return &TicketService{TicketRepository: repo}
}
func (t *TicketService) CreateTicket(vehicleID, spotID, floorID, lotID string, spotType model.SpotType) (*model.Ticket, error) {
	ticket := model.NewTicket(utilies.NewIdGenerator().GenerateID(), vehicleID, spotID, floorID, lotID, time.Now(), spotType)
	if err := t.TicketRepository.SaveTicket(*ticket); err != nil {
		return nil, err
	}
	return ticket, nil
}
func (t *TicketService) GetTicketByID(id string) (*model.Ticket, error) {
	ticket, ok := t.TicketRepository.GetTicketByID(id)
	if !ok {
		return nil, fmt.Errorf("ticket %q not found", id)
	}
	return &ticket, nil
}
func (t *TicketService) CompleteTicket(id string, exit time.Time) error {
	ticket, err := t.GetTicketByID(id)
	if err != nil {
		return err
	}
	if err := ticket.Complete(exit); err != nil {
		return err
	}
	return t.TicketRepository.UpdateTicket(id, *ticket)
}
