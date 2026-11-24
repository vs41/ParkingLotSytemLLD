package main

import (
	"fmt"
	repository "parkinglot/Repository"
	"parkinglot/database"
	"parkinglot/model"
	"parkinglot/service"
	"parkinglot/utilies"
)

func main() {
	db := database.NewDatabase()
	paymentRepo := repository.NewPaymentRepository(db)
	ticketRepo := repository.NewTicketRepository(db)

	floors := make([]*model.Floor, 0, 3)
	spotTypes := []model.SpotType{model.Bike, model.Car, model.Truck}
	for i := 0; i < 3; i++ {
		spots := make([]*model.ParkingSpot, 0, 6)
		for j := 0; j < 6; j++ {
			spots = append(spots, model.NewParkingSpot(utilies.NewIdGenerator().GenerateID(), false, spotTypes[j%len(spotTypes)]))
		}
		floors = append(floors, model.NewFloor(utilies.NewIdGenerator().GenerateID(), model.Active, spots))
	}

	lot := model.NewParkingLot(utilies.NewIdGenerator().GenerateID(), floors, "High Tower")
	lot.PricingType = "HOURLY" // Change to DAILY for this parking lot's pricing policy.

	feeCalculator := service.NewFeeCalculatorService()
	paymentService := service.NewPaymentService(paymentRepo)
	allocationService := service.NewSpotAllocationService()
	ticketService := service.NewTicketService(ticketRepo)
	parkingService := service.NewParkingService(lot, feeCalculator, paymentService, allocationService, ticketService)

	vehicle := model.NewVehicle(utilies.NewIdGenerator().GenerateID(), model.Bike)
	ticket, err := parkingService.ParkVehicle(*vehicle)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Ticket generated: %+v\n", ticket)

	// Customer chooses a payment method at exit time: UPI, CARD, or CASH.
	if err := parkingService.UnParkVehicle(ticket.TicketID, "UPI"); err != nil {
		panic(err)
	}
	fmt.Println("Vehicle exited and payment was recorded successfully.")
}
