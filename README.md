<h1 align="center">
  macos-battery-exporter
</h1>

<p align="center">
  <strong>
    Prometheus exporter for detailed battery metrics on macOS.
  </strong>
</p>

<p align="center">
  <a href="https://github.com/jimeh/macos-battery-exporter/releases"><img src="https://img.shields.io/github/v/tag/jimeh/macos-battery-exporter?label=release" alt="GitHub tag (latest SemVer)"></a>
  <a href="https://pkg.go.dev/github.com/jimeh/macos-battery-exporter"><img src="https://img.shields.io/badge/%E2%80%8B-reference-387b97.svg?logo=go&logoColor=white" alt="Go Reference"></a>
  <a href="https://github.com/jimeh/macos-battery-exporter/issues"><img src="https://img.shields.io/github/issues-raw/jimeh/macos-battery-exporter.svg?style=flat&logo=github&logoColor=white" alt="GitHub issues"></a>
  <a href="https://github.com/jimeh/macos-battery-exporter/pulls"><img src="https://img.shields.io/github/issues-pr-raw/jimeh/macos-battery-exporter.svg?style=flat&logo=github&logoColor=white" alt="GitHub pull requests"></a>
  <a href="https://github.com/jimeh/macos-battery-exporter/blob/main/LICENSE"><img src="https://img.shields.io/github/license/jimeh/macos-battery-exporter.svg?style=flat" alt="License Status"></a>
</p>

A Prometheus exporter for macOS battery information, exposing wide range of
details available from `ioreg` about batteries. Includes a lot more details than
what `node_exporter` supports via its `node_power_supply_*` metrics.

## Installation

### Manually

Binary releases are available on the
[Releases](https://github.com/jimeh/macos-battery-exporter/releases) page.

### Homebrew

You can install it from the [`jimeh/tap`](https://github.com/jimeh/homebrew-tap)
Tap:

```bash
brew install jimeh/tap/macos-battery-exporter
```

#### Service

The homebrew formula has a service configuration, which can be started with:

```bash
brew services start macos-battery-exporter
```

After which battery metrics are available on
[`http://localhost:9108/metrics`](http://localhost:9108/metrics).

### Go

```bash
go install github.com/jimeh/macos-battery-exporter@latest
```

## Usage

### Help

```bash
macos-battery-exporter -h
```

```
usage: macos-battery-exporter [<options>]

  -b string
    	Bind address to run server on (default "127.0.0.1")
  -d string
    	Log output device (stderr or stdout) (default "stderr")
  -l string
    	Log level (default "info")
  -n string
    	Namespace for metrics (default "macos")
  -o string
    	Output file to write to in Prometheus format
  -p int
    	Port to run server on (default 9108)
  -s	Run as a Prometheus metrics server
  -v	Print version and exit
```

### Print to STDOUT

```bash
macos-battery-exporter
```

```ini
# HELP macos_battery_cell_disconnect_count Total number of times a battery cell has been disconnected.
# TYPE macos_battery_cell_disconnect_count gauge
macos_battery_cell_disconnect_count{serial="ZTMDHJEZ8JKMYVAJKU"} 0
# HELP macos_battery_charge_rate_amps Current charge rate in Ah.
# TYPE macos_battery_charge_rate_amps gauge
macos_battery_charge_rate_amps{serial="ZTMDHJEZ8JKMYVAJKU"} -1.004
# HELP macos_battery_charge_rate_watts Current charge rate in Wh.
# TYPE macos_battery_charge_rate_watts gauge
macos_battery_charge_rate_watts{serial="ZTMDHJEZ8JKMYVAJKU"} -12.712648
# HELP macos_battery_count Total number of batteries.
# TYPE macos_battery_count gauge
macos_battery_count 1
# HELP macos_battery_current_capacity_amps Current charge capacity in Ah.
# TYPE macos_battery_current_capacity_amps gauge
macos_battery_current_capacity_amps{serial="ZTMDHJEZ8JKMYVAJKU"} 5.39
# HELP macos_battery_current_capacity_watts Current charge capacity in Wh.
# TYPE macos_battery_current_capacity_watts gauge
macos_battery_current_capacity_watts{serial="ZTMDHJEZ8JKMYVAJKU"} 68.24818
# HELP macos_battery_current_percentage Current battery charge percentage.
# TYPE macos_battery_current_percentage gauge
macos_battery_current_percentage{serial="ZTMDHJEZ8JKMYVAJKU"} 91
# HELP macos_battery_cycle_count Current battery cycle count.
# TYPE macos_battery_cycle_count counter
macos_battery_cycle_count{serial="ZTMDHJEZ8JKMYVAJKU"} 15
# HELP macos_battery_design_capacity_amps Design capacity in Ah.
# TYPE macos_battery_design_capacity_amps gauge
macos_battery_design_capacity_amps{serial="ZTMDHJEZ8JKMYVAJKU"} 6.249
# HELP macos_battery_design_capacity_watts Design capacity in Wh.
# TYPE macos_battery_design_capacity_watts gauge
macos_battery_design_capacity_watts{serial="ZTMDHJEZ8JKMYVAJKU"} 79.124838
# HELP macos_battery_fully_charged Indicates if the battery is fully charged.
# TYPE macos_battery_fully_charged gauge
macos_battery_fully_charged{serial="ZTMDHJEZ8JKMYVAJKU"} 0
# HELP macos_battery_health Battery health as a percentage (0-100%).
# TYPE macos_battery_health gauge
macos_battery_health{serial="ZTMDHJEZ8JKMYVAJKU"} 100
# HELP macos_battery_info Basic details about the battery.
# TYPE macos_battery_info gauge
macos_battery_info{built_in="true",device_name="ayzo3hgs",serial="ZTMDHJEZ8JKMYVAJKU"} 1
# HELP macos_battery_is_charging Indicates if the battery is currently charging.
# TYPE macos_battery_is_charging gauge
macos_battery_is_charging{serial="ZTMDHJEZ8JKMYVAJKU"} 0
# HELP macos_battery_max_capacity_amps Design capacity in Ah.
# TYPE macos_battery_max_capacity_amps gauge
macos_battery_max_capacity_amps{serial="ZTMDHJEZ8JKMYVAJKU"} 6.271
# HELP macos_battery_max_capacity_watts Design capacity in Wh.
# TYPE macos_battery_max_capacity_watts gauge
macos_battery_max_capacity_watts{serial="ZTMDHJEZ8JKMYVAJKU"} 79.403402
# HELP macos_battery_temperature_celsius Current battery temperature in °C.
# TYPE macos_battery_temperature_celsius gauge
macos_battery_temperature_celsius{serial="ZTMDHJEZ8JKMYVAJKU"} 30.53
# HELP macos_battery_time_remaining_seconds Estimated time remaining until battery is fully charged or discharged.
# TYPE macos_battery_time_remaining_seconds gauge
macos_battery_time_remaining_seconds{serial="ZTMDHJEZ8JKMYVAJKU"} 14040
# HELP macos_battery_voltage_volts Current battery voltage in V.
# TYPE macos_battery_voltage_volts gauge
macos_battery_voltage_volts{serial="ZTMDHJEZ8JKMYVAJKU"} 12.662
```

### Write to File

```bash
macos-battery-exporter -o battery.txt
```

### Run Server

```bash
macos-battery-exporter -s
```

```bash
curl http://localhost:9108/metrics
```

## License

[MIT](https://github.com/jimeh/macos-battery-exporter/blob/main/LICENSE)
