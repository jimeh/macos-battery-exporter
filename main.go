//go:build darwin

package main

import (
	"flag"
	"fmt"
	"io"
	"log/slog"
	"os"
	"strings"

	"github.com/jimeh/macos-battery-exporter/prombat"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/expfmt"
)

var (
	name    = "macos-battery-exporter"
	version = "0.0.0-dev"
	commit  = "unknown"
)

type configuration struct {
	Output       string
	Server       bool
	Bind         string
	Port         int
	Namespace    string
	LogLevel     string
	LogDevice    string
	PrintVersion bool
}

func configure() (*configuration, *flag.FlagSet, error) {
	fs := flag.NewFlagSet(name, flag.ExitOnError)
	fs.Usage = func() {
		fmt.Fprintf(fs.Output(),
			"usage: %s [<options>]\n\n", fs.Name(),
		)
		fs.PrintDefaults()
	}

	config := &configuration{}

	fs.StringVar(&config.Output,
		"o", "", "Output file to write to in Prometheus format",
	)
	fs.BoolVar(&config.Server,
		"s", false, "Run as a Prometheus metrics server",
	)
	fs.StringVar(&config.Bind,
		"b", "127.0.0.1", "Bind address to run server on",
	)
	fs.IntVar(&config.Port, "p", 9108, "Port to run server on")
	fs.StringVar(&config.Namespace,
		"n", prombat.DefaultNamespace, "Namespace for metrics",
	)
	fs.StringVar(&config.LogLevel,
		"l", "info", "Log level",
	)
	fs.StringVar(&config.LogDevice,
		"d", "stderr", "Log output device (stderr or stdout)",
	)
	fs.BoolVar(&config.PrintVersion,
		"v", false, "Print version and exit",
	)

	err := fs.Parse(os.Args[1:])
	if err != nil {
		return nil, nil, err
	}

	return config, fs, nil
}

func main() {
	if err := mainE(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}

func mainE() error {
	config, _, err := configure()
	if err != nil {
		return err
	}

	err = setupSLog(config.LogDevice, config.LogLevel)
	if err != nil {
		return err
	}

	if config.PrintVersion {
		fmt.Printf("%s %s (%s)\n", name, version, commit)
		return nil
	}

	if config.Server {
		return runServer(config)
	}

	metrics, err := renderMetrics(config)
	if err != nil {
		return err
	}

	if config.Output != "" {
		return writeToFile(metrics, config.Output)
	}

	fmt.Print(metrics)
	return nil
}

func runServer(config *configuration) error {
	opts := prombat.ServerOptions{
		Bind: config.Bind,
		Port: config.Port,
	}

	return prombat.RunServer(
		config.Namespace,
		prometheus.DefaultRegisterer.(*prometheus.Registry),
		opts,
	)
}

func renderMetrics(config *configuration) (string, error) {
	registry := prometheus.NewRegistry()
	err := registry.Register(prombat.NewCollector(config.Namespace))
	if err != nil {
		return "", err
	}

	gatherers := prometheus.Gatherers{registry}
	metricFamilies, err := gatherers.Gather()
	if err != nil {
		return "", err
	}

	var sb strings.Builder
	for _, mf := range metricFamilies {
		_, err := expfmt.MetricFamilyToText(&sb, mf)
		if err != nil {
			return "", err
		}
	}

	return sb.String(), nil
}

func setupSLog(device string, levelStr string) error {
	var w io.Writer
	switch device {
	case "stderr":
		w = os.Stderr
	case "stdout":
		w = os.Stdout
	default:
		return fmt.Errorf("invalid log device: %s", device)
	}

	var level slog.Level
	err := level.UnmarshalText([]byte(levelStr))
	if err != nil {
		return err
	}

	handler := slog.NewTextHandler(w, &slog.HandlerOptions{
		Level: level,
	})
	logger := slog.New(handler)

	slog.SetDefault(logger)

	return nil
}

func writeToFile(data, outputFile string) error {
	file, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(data)
	if err != nil {
		return err
	}

	return nil
}
