//go:build darwin

package battery

import (
	"bytes"
	"os/exec"

	"howett.net/plist"
)

// batteryRaw is the raw data structure returned by ioreg.
type batteryRaw struct {
	Amperage                   int64  `plist:"Amperage"`
	AvgTimeToEmpty             int    `plist:"AvgTimeToEmpty"`
	AvgTimeToFull              int    `plist:"AvgTimeToFull"`
	BatteryCellDisconnectCount int    `plist:"BatteryCellDisconnectCount"`
	BuiltIn                    bool   `plist:"built-in"`
	CurrentCapacity            int    `plist:"AppleRawCurrentCapacity"`
	CurrentPercentage          int    `plist:"CurrentCapacity"`
	CycleCount                 int    `plist:"CycleCount"`
	DesignCapacity             int    `plist:"DesignCapacity"`
	DeviceName                 string `plist:"DeviceName"`
	DesignCycleCount           int    `plist:"DesignCycleCount9C"`
	ExternalConnected          bool   `plist:"ExternalConnected"`
	FullyCharged               bool   `plist:"FullyCharged"`
	Health                     int    `plist:"MaxCapacity"`
	IsCharging                 bool   `plist:"IsCharging"`
	MaxCapacity                int    `plist:"AppleRawMaxCapacity"`
	Serial                     string `plist:"Serial"`
	Temperature                int    `plist:"Temperature"`
	TimeRemaining              int    `plist:"TimeRemaining"`
	Voltage                    int64  `plist:"Voltage"`
}

func getAllRaw() ([]*batteryRaw, error) {
	b, err := exec.Command("ioreg", "-ra", "-c", "AppleSmartBattery").Output()
	if err != nil {
		return nil, err
	}

	batteries := []*batteryRaw{}

	if len(bytes.TrimSpace(b)) == 0 {
		return batteries, nil
	}

	_, err = plist.Unmarshal(b, &batteries)
	if err != nil {
		return nil, err
	}

	return batteries, nil
}
