package model

import "testing"

func TestParkingSpotStateTransitions(t *testing.T) {
	spot := NewParkingSpot("S1", false, Car)
	if got := spot.GetStatus(); got != SpotAvailable {
		t.Fatalf("initial status = %q, want %q", got, SpotAvailable)
	}
	if !spot.AssignVehicle("V1") {
		t.Fatal("available spot should accept a vehicle")
	}
	if got := spot.GetStatus(); got != SpotOccupied {
		t.Fatalf("status after assignment = %q, want %q", got, SpotOccupied)
	}
	if spot.AssignVehicle("V2") {
		t.Fatal("occupied spot must reject another vehicle")
	}
	if !spot.RemoveVehicle() {
		t.Fatal("occupied spot should allow vehicle removal")
	}
	if got := spot.GetStatus(); got != SpotAvailable {
		t.Fatalf("status after removal = %q, want %q", got, SpotAvailable)
	}
}

func TestParkingSpotOutOfServiceTransitions(t *testing.T) {
	spot := NewParkingSpot("S2", false, Bike)
	if !spot.SetOutOfService() {
		t.Fatal("available spot should be taken out of service")
	}
	if spot.AssignVehicle("V1") {
		t.Fatal("out-of-service spot must reject allocation")
	}
	if spot.RemoveVehicle() {
		t.Fatal("out-of-service spot cannot remove a vehicle")
	}
	if !spot.RestoreToService() {
		t.Fatal("out-of-service spot should be restorable")
	}
	if !spot.IsAvailable() {
		t.Fatal("restored spot should be available")
	}
}
