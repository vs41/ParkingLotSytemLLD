package service

import (
	"fmt"
	"parkinglot/model"
)

type SpotAllocationService struct{}

func NewSpotAllocationService() *SpotAllocationService { return &SpotAllocationService{} }
func (s *SpotAllocationService) FindAvailableSpot(floors []*model.Floor, vehicleType model.SpotType) (*model.ParkingSpot, string) {
	for _, floor := range floors {
		for _, spot := range floor.ParkingSpots {
			if spot.IsAvailable() && spot.SpotTypes == vehicleType {
				return spot, floor.FloorID
			}
		}
	}
	return nil, ""
}
func (s *SpotAllocationService) AllocateSpot(floors []*model.Floor, vehicle *model.Vehicle) (*model.ParkingSpot, string, error) {
	if vehicle == nil {
		return nil, "", fmt.Errorf("vehicle is required")
	}
	spot, floorID := s.FindAvailableSpot(floors, vehicle.VehicleType)
	if spot == nil {
		return nil, "", fmt.Errorf("no available spot for vehicle type %s", vehicle.VehicleType)
	}
	if !spot.AssignVehicle(vehicle.VehicleID) {
		return nil, "", fmt.Errorf("spot %s is no longer available", spot.SpotID)
	}
	return spot, floorID, nil
}
func (s *SpotAllocationService) ReleaseSpot(floors []*model.Floor, spotID, floorID string) error {
	for _, floor := range floors {
		for _, spot := range floor.ParkingSpots {
			if spot.SpotID == spotID && floor.FloorID == floorID {
				if spot.IsAvailable() {
					return fmt.Errorf("spot %s is already available", spotID)
				}
				spot.RemoveVehicle()
				return nil
			}
		}
	}
	return fmt.Errorf("spot %s not found on floor %s", spotID, floorID)
}
