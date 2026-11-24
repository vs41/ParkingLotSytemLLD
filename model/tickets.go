package model

import (
	"fmt"
	"time"
)

type TicketStatus string

const (
	TicketActive   TicketStatus = "Active"
	TicketInactive TicketStatus = "Inactive"
)

type Ticket struct {
	TicketID     string
	VehicleID    string
	SpotID       string
	FloorID      string
	SpotType     SpotType
	ParkingLotID string
	Status       TicketStatus
	ParkingTime  time.Time
	ExitTime     time.Time
}

func NewTicket(ticketID, vehicleID, spotID, floorID, parkingLotID string, parkingTime time.Time, spotType SpotType) *Ticket {
	return &Ticket{TicketID: ticketID, VehicleID: vehicleID, SpotID: spotID, FloorID: floorID, SpotType: spotType, ParkingLotID: parkingLotID, Status: TicketActive, ParkingTime: parkingTime}
}

func (t *Ticket) GetTicketID() string     { return t.TicketID }
func (t *Ticket) GetVehicleID() string    { return t.VehicleID }
func (t *Ticket) GetSpotID() string       { return t.SpotID }
func (t *Ticket) GetFloorID() string      { return t.FloorID }
func (t *Ticket) GetParkingLotID() string { return t.ParkingLotID }
func (t *Ticket) GetStatus() TicketStatus { return t.Status }

// Complete performs the ticket's only valid lifecycle transition.
func (t *Ticket) Complete(exitTime time.Time) error {
	if t == nil {
		return fmt.Errorf("ticket is required")
	}
	if t.Status != TicketActive {
		return fmt.Errorf("ticket %q is not active", t.TicketID)
	}
	if exitTime.Before(t.ParkingTime) {
		return fmt.Errorf("exit time cannot be before parking time")
	}
	t.ExitTime = exitTime
	t.Status = TicketInactive
	return nil
}
