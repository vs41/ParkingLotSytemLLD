package model

type ParkingLot struct {
	ParkingLotID string
	Floors       []*Floor
	Name         string
	PricingType  string
}

func NewParkingLot(id string, floors []*Floor, name string, _ ...int) *ParkingLot {
	return &ParkingLot{ParkingLotID: id, Floors: floors, Name: name, PricingType: "HOURLY"}
}
func (pl *ParkingLot) GetParkingLotID() string { return pl.ParkingLotID }
func (pl *ParkingLot) GetFloors() []*Floor     { return pl.Floors }
func (pl *ParkingLot) GetName() string         { return pl.Name }
func (pl *ParkingLot) AddFloor(floor *Floor)   { pl.Floors = append(pl.Floors, floor) }
func (pl *ParkingLot) RemoveFloor(id string) bool {
	for i, floor := range pl.Floors {
		if floor.FloorID == id {
			pl.Floors = append(pl.Floors[:i], pl.Floors[i+1:]...)
			return true
		}
	}
	return false
}
