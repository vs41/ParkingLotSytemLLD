package model

// SpotType describes which vehicles a spot can accommodate.
type SpotType string

const (
	Car   SpotType = "Car"
	Bike  SpotType = "Bike"
	Truck SpotType = "Truck"
)

// SpotStatus is the externally visible lifecycle state of a parking spot.
type SpotStatus string

const (
	SpotAvailable    SpotStatus = "AVAILABLE"
	SpotOccupied     SpotStatus = "OCCUPIED"
	SpotOutOfService SpotStatus = "OUT_OF_SERVICE"
)

// ParkingSpotState encapsulates behavior that changes depending on a spot's state.
// The state implementations are private; callers interact through ParkingSpot.
type ParkingSpotState interface {
	AssignVehicle(spot *ParkingSpot, vehicleID string) bool
	RemoveVehicle(spot *ParkingSpot) bool
	Status() SpotStatus
}

type availableSpotState struct{}
type occupiedSpotState struct{}
type outOfServiceSpotState struct{}

func (availableSpotState) AssignVehicle(spot *ParkingSpot, vehicleID string) bool {
	if vehicleID == "" {
		return false
	}
	spot.VehicleID = vehicleID
	spot.state = occupiedSpotState{}
	return true
}
func (availableSpotState) RemoveVehicle(*ParkingSpot) bool { return false }
func (availableSpotState) Status() SpotStatus              { return SpotAvailable }

func (occupiedSpotState) AssignVehicle(*ParkingSpot, string) bool { return false }
func (occupiedSpotState) RemoveVehicle(spot *ParkingSpot) bool {
	spot.VehicleID = ""
	spot.state = availableSpotState{}
	return true
}
func (occupiedSpotState) Status() SpotStatus { return SpotOccupied }

func (outOfServiceSpotState) AssignVehicle(*ParkingSpot, string) bool { return false }
func (outOfServiceSpotState) RemoveVehicle(*ParkingSpot) bool         { return false }
func (outOfServiceSpotState) Status() SpotStatus                      { return SpotOutOfService }

type ParkingSpot struct {
	SpotID    string
	SpotTypes SpotType
	VehicleID string
	state     ParkingSpotState
}

// NewParkingSpot keeps the existing constructor shape for the project.
// A spot created with occupied=true has no assigned vehicle ID, so new spots
// should normally be created with occupied=false.
func NewParkingSpot(spotID string, occupied bool, spotTypes SpotType) *ParkingSpot {
	spot := &ParkingSpot{SpotID: spotID, SpotTypes: spotTypes, state: availableSpotState{}}
	if occupied {
		spot.state = occupiedSpotState{}
	}
	return spot
}

func (ps *ParkingSpot) AssignVehicle(vehicleID string) bool {
	if ps == nil || ps.state == nil {
		return false
	}
	return ps.state.AssignVehicle(ps, vehicleID)
}

func (ps *ParkingSpot) RemoveVehicle() bool {
	if ps == nil || ps.state == nil {
		return false
	}
	return ps.state.RemoveVehicle(ps)
}

func (ps *ParkingSpot) IsAvailable() bool {
	return ps != nil && ps.state != nil && ps.state.Status() == SpotAvailable
}

func (ps *ParkingSpot) GetStatus() SpotStatus {
	if ps == nil || ps.state == nil {
		return SpotOutOfService
	}
	return ps.state.Status()
}

// SetOutOfService prevents allocation of an available spot. Occupied spots
// cannot be taken out of service until their vehicle has been removed.
func (ps *ParkingSpot) SetOutOfService() bool {
	if ps == nil || ps.state == nil || ps.state.Status() != SpotAvailable {
		return false
	}
	ps.state = outOfServiceSpotState{}
	return true
}

func (ps *ParkingSpot) RestoreToService() bool {
	if ps == nil || ps.state == nil || ps.state.Status() != SpotOutOfService {
		return false
	}
	ps.state = availableSpotState{}
	return true
}

func (ps *ParkingSpot) GetSpotID() string { return ps.SpotID }
