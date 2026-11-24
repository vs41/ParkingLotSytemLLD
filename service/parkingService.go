package service

import (
	"fmt"
	"parkinglot/model"
	"time"
)

type ParkingService struct {
	ParkingLot     *model.ParkingLot
	FeeCalculator  *FeeCalculatorService
	Payment        *PaymentService
	SpotAllocation *SpotAllocationService
	TicketService  *TicketService
}

func NewParkingService(lot *model.ParkingLot, fees *FeeCalculatorService, payments *PaymentService, allocation *SpotAllocationService, tickets *TicketService) *ParkingService {
	return &ParkingService{ParkingLot: lot, FeeCalculator: fees, Payment: payments, SpotAllocation: allocation, TicketService: tickets}
}
func (p *ParkingService) ParkVehicle(vehicle model.Vehicle) (*model.Ticket, error) {
	if p.ParkingLot == nil {
		return nil, fmt.Errorf("parking lot is not configured")
	}
	spot, floorID, err := p.SpotAllocation.AllocateSpot(p.ParkingLot.Floors, &vehicle)
	if err != nil {
		return nil, err
	}
	ticket, err := p.TicketService.CreateTicket(vehicle.VehicleID, spot.SpotID, floorID, p.ParkingLot.ParkingLotID, spot.SpotTypes)
	if err != nil {
		_ = p.SpotAllocation.ReleaseSpot(p.ParkingLot.Floors, spot.SpotID, floorID)
		return nil, err
	}
	return ticket, nil
}
func (p *ParkingService) UnParkVehicle(ticketID, paymentMethod string) error {
	ticket, err := p.TicketService.GetTicketByID(ticketID)
	if err != nil {
		return err
	}
	if ticket.Status != model.TicketActive {
		return fmt.Errorf("ticket %s is not active", ticketID)
	}
	exitTime := time.Now()
	durationHours := exitTime.Sub(ticket.ParkingTime).Hours()
	amount, err := p.FeeCalculator.CalculateFee(p.ParkingLot.PricingType, string(ticket.SpotType), durationHours)
	if err != nil {
		return err
	}
	if err = p.Payment.ProcessPayment(paymentMethod, amount, ticket.TicketID); err != nil {
		return err
	}
	if err = p.TicketService.CompleteTicket(ticket.TicketID, exitTime); err != nil {
		return err
	}
	return p.SpotAllocation.ReleaseSpot(p.ParkingLot.Floors, ticket.SpotID, ticket.FloorID)
}
