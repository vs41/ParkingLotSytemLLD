package model

type Vehicle struct {
	VehicleID   string
	VehicleType SpotType
}

func NewVehicle(vehicleID string, vehicleType SpotType) *Vehicle {
	return &Vehicle{
		VehicleID:   vehicleID,
		VehicleType: vehicleType,
	}
}

func (v *Vehicle) GetVehicleID() string {
	return v.VehicleID
}

func (v *Vehicle) GetVehicleType() SpotType {
	return v.VehicleType
}
