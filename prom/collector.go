//go:build darwin

package prom

import (
	"log/slog"
	"math"

	"github.com/jimeh/macos-battery-exporter/battery"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	serialLabel     = "serial"
	deviceNameLabel = "device_name"
	builtInLabel    = "built_in"
)

type Collector struct {
	descInfo         *prometheus.Desc
	descBatteryCount *prometheus.Desc

	descBatteryCellDisconnectCount *prometheus.Desc
	descChargeRateAmps             *prometheus.Desc
	descChargeRateWatts            *prometheus.Desc
	descCurrentCapacityAmps        *prometheus.Desc
	descCurrentCapacityWatts       *prometheus.Desc
	descCurrentPercentage          *prometheus.Desc
	descCycleCount                 *prometheus.Desc
	descDesignCapacityAmps         *prometheus.Desc
	descDesignCapacityWatts        *prometheus.Desc
	descFullyCharged               *prometheus.Desc
	descHealth                     *prometheus.Desc
	descIsCharging                 *prometheus.Desc
	descMaxCapacityAmps            *prometheus.Desc
	descMaxCapacityWatts           *prometheus.Desc
	descTemperature                *prometheus.Desc
	descTimeRemaining              *prometheus.Desc
	descVoltage                    *prometheus.Desc
}

var _ prometheus.Collector = &Collector{}

func NewCollector(namespace string) *Collector {
	c := &Collector{
		descInfo: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "battery", "info"),
			"Basic details about the battery.",
			[]string{serialLabel, deviceNameLabel, builtInLabel},
			nil,
		),
		descBatteryCount: prometheus.NewDesc(
			prometheus.BuildFQName(
				namespace, "battery", "count",
			),
			"Total number of batteries.",
			nil,
			nil,
		),
		descBatteryCellDisconnectCount: prometheus.NewDesc(
			prometheus.BuildFQName(
				namespace, "battery", "cell_disconnect_count",
			),
			"Total number of times a battery cell has been disconnected.",
			[]string{serialLabel},
			nil,
		),
		descChargeRateAmps: prometheus.NewDesc(
			prometheus.BuildFQName(
				namespace, "battery", "charge_rate_amps",
			),
			"Current charge rate in Ah.",
			[]string{serialLabel},
			nil,
		),
		descChargeRateWatts: prometheus.NewDesc(
			prometheus.BuildFQName(
				namespace, "battery", "charge_rate_watts",
			),
			"Current charge rate in Wh.",
			[]string{serialLabel},
			nil,
		),
		descCurrentCapacityAmps: prometheus.NewDesc(
			prometheus.BuildFQName(
				namespace, "battery", "current_capacity_amps",
			),
			"Current charge capacity in Ah.",
			[]string{serialLabel},
			nil,
		),
		descCurrentCapacityWatts: prometheus.NewDesc(
			prometheus.BuildFQName(
				namespace, "battery", "current_capacity_watts",
			),
			"Current charge capacity in Wh.",
			[]string{serialLabel},
			nil,
		),
		descCurrentPercentage: prometheus.NewDesc(
			prometheus.BuildFQName(
				namespace, "battery", "current_percentage",
			),
			"Current battery charge percentage.",
			[]string{serialLabel},
			nil,
		),
		descCycleCount: prometheus.NewDesc(
			prometheus.BuildFQName(
				namespace, "battery", "cycle_count",
			),
			"Current battery cycle count.",
			[]string{serialLabel},
			nil,
		),
		descDesignCapacityAmps: prometheus.NewDesc(
			prometheus.BuildFQName(
				namespace, "battery", "design_capacity_amps",
			),
			"Design capacity in Ah.",
			[]string{serialLabel},
			nil,
		),
		descDesignCapacityWatts: prometheus.NewDesc(
			prometheus.BuildFQName(
				namespace, "battery", "design_capacity_watts",
			),
			"Design capacity in Wh.",
			[]string{serialLabel},
			nil,
		),
		descFullyCharged: prometheus.NewDesc(
			prometheus.BuildFQName(
				namespace, "battery", "fully_charged",
			),
			"Indicates if the battery is fully charged.",
			[]string{serialLabel},
			nil,
		),
		descHealth: prometheus.NewDesc(
			prometheus.BuildFQName(
				namespace, "battery", "health",
			),
			"Battery health as a percentage (0-100%).",
			[]string{serialLabel},
			nil,
		),
		descIsCharging: prometheus.NewDesc(
			prometheus.BuildFQName(
				namespace, "battery", "is_charging",
			),
			"Indicates if the battery is currently charging.",
			[]string{serialLabel},
			nil,
		),
		descMaxCapacityAmps: prometheus.NewDesc(
			prometheus.BuildFQName(
				namespace, "battery", "max_capacity_amps",
			),
			"Design capacity in Ah.",
			[]string{serialLabel},
			nil,
		),
		descMaxCapacityWatts: prometheus.NewDesc(
			prometheus.BuildFQName(
				namespace, "battery", "max_capacity_watts",
			),
			"Design capacity in Wh.",
			[]string{serialLabel},
			nil,
		),
		descTemperature: prometheus.NewDesc(
			prometheus.BuildFQName(
				namespace, "battery", "temperature_celsius",
			),
			"Current battery temperature in °C.",
			[]string{serialLabel},
			nil,
		),
		descTimeRemaining: prometheus.NewDesc(
			prometheus.BuildFQName(
				namespace, "battery", "time_remaining_seconds",
			),
			"Estimated time remaining until battery is fully "+
				"charged or discharged.",
			[]string{serialLabel},
			nil,
		),
		descVoltage: prometheus.NewDesc(
			prometheus.BuildFQName(
				namespace, "battery", "voltage_volts",
			),
			"Current battery voltage in V.",
			[]string{serialLabel},
			nil,
		),
	}

	return c
}

func (c *Collector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.descInfo
	ch <- c.descBatteryCount
	ch <- c.descBatteryCellDisconnectCount
	ch <- c.descChargeRateAmps
	ch <- c.descChargeRateWatts
	ch <- c.descCurrentCapacityAmps
	ch <- c.descCurrentCapacityWatts
	ch <- c.descCurrentPercentage
	ch <- c.descCycleCount
	ch <- c.descDesignCapacityAmps
	ch <- c.descDesignCapacityWatts
	ch <- c.descFullyCharged
	ch <- c.descHealth
	ch <- c.descIsCharging
	ch <- c.descMaxCapacityAmps
	ch <- c.descMaxCapacityWatts
	ch <- c.descTemperature
	ch <- c.descTimeRemaining
}

func (c *Collector) Collect(ch chan<- prometheus.Metric) {
	slog.Debug("collecting battery metrics")
	batteries, err := battery.GetAll()
	if err != nil {
		slog.Error(
			"failed to get battery details",
			slog.String("error", err.Error()),
		)
		return
	}

	ch <- prometheus.MustNewConstMetric(
		c.descBatteryCount,
		prometheus.GaugeValue,
		float64(len(batteries)),
	)

	for _, battery := range batteries {
		labels := []string{battery.Serial}

		ch <- prometheus.MustNewConstMetric(
			c.descInfo,
			prometheus.GaugeValue,
			1,
			battery.Serial, battery.DeviceName, boolToString(battery.BuiltIn),
		)
		ch <- prometheus.MustNewConstMetric(
			c.descBatteryCellDisconnectCount,
			prometheus.GaugeValue,
			float64(battery.BatteryCellDisconnectCount),
			labels...,
		)
		ch <- prometheus.MustNewConstMetric(
			c.descChargeRateAmps,
			prometheus.GaugeValue,
			roundTo(float64(battery.ChargeRateAmps)/1000, 6),
			labels...,
		)
		ch <- prometheus.MustNewConstMetric(
			c.descChargeRateWatts,
			prometheus.GaugeValue,
			roundTo(battery.ChargeRateWatts/1000, 6),
			labels...,
		)
		ch <- prometheus.MustNewConstMetric(
			c.descCurrentCapacityAmps,
			prometheus.GaugeValue,
			roundTo(float64(battery.CurrentCapacityAmps)/1000, 6),
			labels...,
		)
		ch <- prometheus.MustNewConstMetric(
			c.descCurrentCapacityWatts,
			prometheus.GaugeValue,
			roundTo(battery.CurrentCapacityWatts/1000, 6),
			labels...,
		)
		ch <- prometheus.MustNewConstMetric(
			c.descCurrentPercentage,
			prometheus.GaugeValue,
			float64(battery.CurrentPercentage),
			labels...,
		)
		ch <- prometheus.MustNewConstMetric(
			c.descCycleCount,
			prometheus.CounterValue,
			float64(battery.CycleCount),
			labels...,
		)
		ch <- prometheus.MustNewConstMetric(
			c.descDesignCapacityAmps,
			prometheus.GaugeValue,
			roundTo(float64(battery.DesignCapacityAmps)/1000, 6),
			labels...,
		)
		ch <- prometheus.MustNewConstMetric(
			c.descDesignCapacityWatts,
			prometheus.GaugeValue,
			roundTo(battery.DesignCapacityWatts/1000, 6),
			labels...,
		)
		ch <- prometheus.MustNewConstMetric(
			c.descFullyCharged,
			prometheus.GaugeValue,
			boolToFloat64(battery.FullyCharged),
			labels...,
		)
		ch <- prometheus.MustNewConstMetric(
			c.descHealth,
			prometheus.GaugeValue,
			float64(battery.Health),
			labels...,
		)
		ch <- prometheus.MustNewConstMetric(
			c.descIsCharging,
			prometheus.GaugeValue,
			boolToFloat64(battery.IsCharging),
			labels...,
		)
		ch <- prometheus.MustNewConstMetric(
			c.descMaxCapacityAmps,
			prometheus.GaugeValue,
			roundTo(float64(battery.MaxCapacityAmps)/1000, 6),
			labels...,
		)
		ch <- prometheus.MustNewConstMetric(
			c.descMaxCapacityWatts,
			prometheus.GaugeValue,
			roundTo(battery.MaxCapacityWatts/1000, 6),
			labels...,
		)
		ch <- prometheus.MustNewConstMetric(
			c.descTemperature,
			prometheus.GaugeValue,
			battery.Temperature,
			labels...,
		)
		ch <- prometheus.MustNewConstMetric(
			c.descTimeRemaining,
			prometheus.GaugeValue,
			battery.TimeRemaining.Seconds(),
			labels...,
		)
		ch <- prometheus.MustNewConstMetric(
			c.descVoltage,
			prometheus.GaugeValue,
			roundTo(float64(battery.Voltage)/1000, 6),
			labels...,
		)
	}
}

func boolToString(b bool) string {
	if b {
		return "true"
	}

	return "false"
}

func boolToFloat64(b bool) float64 {
	if b {
		return 1
	}

	return 0
}

// roundTo rounds a float64 to 'places' decimal places.
//
//nolint:unparam
func roundTo(value float64, places int) float64 {
	shift := math.Pow(10, float64(places))
	return math.Round(value*shift) / shift
}
