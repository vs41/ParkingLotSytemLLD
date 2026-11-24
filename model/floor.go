package model

type floorStatus string

const (
	Active   floorStatus = "Active"
	Inactive floorStatus = "Inactive"
	Full     floorStatus = "Full"
)

type Floor struct {
	FloorID      string
	FloorStatus  floorStatus
	ParkingSpots []*ParkingSpot
}

func NewFloor(floorID string, floorStatus floorStatus, parkingSpots []*ParkingSpot) *Floor {
	return &Floor{
		FloorID:      floorID,
		FloorStatus:  floorStatus,
		ParkingSpots: parkingSpots,
	}
}

func (f *Floor) GetFloorID() string {
	return f.FloorID
}

func (f *Floor) GetFloorStatus() floorStatus {
	return f.FloorStatus
}

func (f *Floor) AvailableParkingSpots() *ParkingSpot {
	for _, spot := range f.ParkingSpots {
		if spot.IsAvailable() == true {
			return spot
		}
	}
	return nil
}

func (f *Floor) AddParkingSpot(spot *ParkingSpot) {
	f.ParkingSpots = append(f.ParkingSpots, spot)
}

func (f *Floor) RemoveParkingSpot(spotId string) bool {
	for index, spot := range f.ParkingSpots {
		if spot.GetSpotID() == spotId {
			f.ParkingSpots = append(f.ParkingSpots[:index], f.ParkingSpots[index+1:]...)
			return true
		}
	}
	return false
}
