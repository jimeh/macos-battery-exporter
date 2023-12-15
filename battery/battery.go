//go:build darwin

package battery

import (
	"math"
	"time"
)

type Battery struct {
	// BatteryCellDisconnectCount is the number of times the battery cells have
	// been disconnected.
	BatteryCellDisconnectCount int

	// BuiltIn indicates if the battery is built-in or not.
	BuiltIn bool

	// ChargeRateAmps is the current charge rate in mAh. Negative values indicate
	// discharge, positive values indicate charging.
	ChargeRateAmps int64

	// ChargeRateWatts is the current charge rate in mWh. Negative values indicate
	// discharge, positive values indicate charging.
	ChargeRateWatts float64

	// CurrentCapacityAmps is the current battery capacity in mAh.
	CurrentCapacityAmps int

	// CurrentCapacityWatts is the current battery capacity in mWh.
	CurrentCapacityWatts float64

	// CurrentPercentage is the current battery capacity as a percentage.
	CurrentPercentage int

	// CycleCount is the current cycle count.
	CycleCount int

	// DesignCapacityAmps is the design capacity in mAh.
	DesignCapacityAmps int

	// DesignCapacityWatts is the design capacity in mWh.
	DesignCapacityWatts float64

	// DeviceName is the battery device name.
	DeviceName string

	// FullyCharged indicates if the battery is fully charged.
	FullyCharged bool

	// Health is the battery health as a percentage (0-100%).
	Health int

	// IsCharging indicates if the battery is currently charging.
	IsCharging bool

	// MaxCapacityAmps is the maximum capacity in mAh.
	MaxCapacityAmps int

	// MaxCapacityWatts is the maximum capacity in mWh.
	MaxCapacityWatts float64

	// Serial is the battery serial number.
	Serial string

	// Temperature is the current temperature in °C.
	Temperature float64

	// TimeRemaining is the estimated time remaining until the battery is
	// fully charged or discharged.
	TimeRemaining time.Duration

	// Voltage is the current voltage in mV.
	Voltage int64
}

func newBattery(b *batteryRaw) *Battery {
	volts := float64(b.Voltage) / 1000

	return &Battery{
		BatteryCellDisconnectCount: b.BatteryCellDisconnectCount,
		BuiltIn:                    b.BuiltIn,
		ChargeRateAmps:             b.Amperage,
		ChargeRateWatts:            roundTo(float64(b.Amperage)*volts, 3),
		CurrentCapacityAmps:        b.CurrentCapacity,
		CurrentCapacityWatts:       roundTo(float64(b.CurrentCapacity)*volts, 3),
		CurrentPercentage:          b.CurrentPercentage,
		CycleCount:                 b.CycleCount,
		DesignCapacityAmps:         b.DesignCapacity,
		DesignCapacityWatts:        roundTo(float64(b.DesignCapacity)*volts, 3),
		DeviceName:                 b.DeviceName,
		FullyCharged:               b.FullyCharged,
		Health:                     b.Health,
		IsCharging:                 b.IsCharging,
		MaxCapacityAmps:            b.MaxCapacity,
		MaxCapacityWatts:           roundTo(float64(b.MaxCapacity)*volts, 3),
		Serial:                     b.Serial,
		Temperature:                float64(b.Temperature) / 100,
		TimeRemaining:              time.Duration(b.TimeRemaining) * time.Minute,
		Voltage:                    b.Voltage,
	}
}

func Get() (*Battery, error) {
	batteriesRaw, err := getAllRaw()
	if err != nil {
		return nil, err
	}

	return newBattery(batteriesRaw[0]), nil
}

func GetAll() ([]*Battery, error) {
	batteriesRaw, err := getAllRaw()
	if err != nil {
		return nil, err
	}

	batteries := []*Battery{}
	for _, b := range batteriesRaw {
		batteries = append(batteries, newBattery(b))
	}

	return batteries, nil
}

// roundTo rounds a float64 to 'places' decimal places
func roundTo(value float64, places int) float64 {
	shift := math.Pow(10, float64(places))
	return math.Round(value*shift) / shift
}
