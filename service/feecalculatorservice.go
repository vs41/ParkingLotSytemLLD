package service

import (
	"fmt"
	"strings"
)

type FeeCalculationStrategy interface {
	CalculateFee(vehicleType string, durationHours float64) (float64, error)
}

type FeeCalculatorPerHours struct{ CarRate, BikeRate, TruckRate float64 }
type FeeCalculatorPerDays struct{ CarRate, BikeRate, TruckRate float64 }

func NewHourlyFeeCalculator() *FeeCalculatorPerHours {
	return &FeeCalculatorPerHours{CarRate: 10, BikeRate: 5, TruckRate: 15}
}
func NewDailyFeeCalculator() *FeeCalculatorPerDays {
	return &FeeCalculatorPerDays{CarRate: 100, BikeRate: 50, TruckRate: 150}
}
func feeFor(vehicle string, car, bike, truck float64) (float64, error) {
	switch vehicle {
	case "Car":
		return car, nil
	case "Bike":
		return bike, nil
	case "Truck":
		return truck, nil
	default:
		return 0, fmt.Errorf("unsupported vehicle type %q", vehicle)
	}
}
func (f *FeeCalculatorPerHours) CalculateFee(vehicle string, hours float64) (float64, error) {
	if hours < 0 {
		return 0, fmt.Errorf("duration cannot be negative")
	}
	rate, err := feeFor(vehicle, f.CarRate, f.BikeRate, f.TruckRate)
	return rate * hours, err
}
func (f *FeeCalculatorPerDays) CalculateFee(vehicle string, hours float64) (float64, error) {
	if hours < 0 {
		return 0, fmt.Errorf("duration cannot be negative")
	}
	rate, err := feeFor(vehicle, f.CarRate, f.BikeRate, f.TruckRate)
	return rate * (hours / 24), err
}

// FeeCalculatorService selects the configured pricing strategy for a lot.
type FeeCalculatorService struct {
	strategies map[string]FeeCalculationStrategy
}

func NewFeeCalculatorService() *FeeCalculatorService {
	return &FeeCalculatorService{strategies: map[string]FeeCalculationStrategy{"HOURLY": NewHourlyFeeCalculator(), "DAILY": NewDailyFeeCalculator()}}
}
func (f *FeeCalculatorService) RegisterStrategy(name string, strategy FeeCalculationStrategy) error {
	key := strings.ToUpper(strings.TrimSpace(name))
	if key == "" || strategy == nil {
		return fmt.Errorf("pricing name and strategy are required")
	}
	f.strategies[key] = strategy
	return nil
}
func (f *FeeCalculatorService) CalculateFee(pricingType, vehicleType string, durationHours float64) (float64, error) {
	strategy, ok := f.strategies[strings.ToUpper(strings.TrimSpace(pricingType))]
	if !ok {
		return 0, fmt.Errorf("unsupported pricing type %q", pricingType)
	}
	return strategy.CalculateFee(vehicleType, durationHours)
}
