//go:build darwin

package main

import (
	"flag"
	"fmt"
	"log"
	"log/slog"
	"os"
	"strings"

	"github.com/jimeh/macos-battery-exporter/prombat"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/expfmt"
)

var (
	version = "0.0.0-dev"
	commit  = "unknown"

	outputFlag = flag.String(
		"o", "", "Output file to write to in Prometheus format",
	)
	serverFlag = flag.Bool("s", false, "Run as a Prometheus metrics server")
	bindFlag   = flag.String(
		"b", "127.0.0.1", "Bind address to run server on",
	)
	portFlag      = flag.Int("p", 9108, "Port to run server on")
	namespaceFlag = flag.String(
		"n", prombat.DefaultNamespace, "Namespace for metrics",
	)
	logLevelFlag = flag.String("l", "info", "Log level")
	versionFlag  = flag.Bool("v", false, "Print version and exit")
)

func main() {
	if err := mainE(); err != nil {
		log.Fatal(err)
	}
}

func mainE() error {
	flag.Parse()

	err := setupSLog(*logLevelFlag)
	if err != nil {
		return err
	}

	if *versionFlag {
		fmt.Printf("macos-battery-exporter %s (%s)\n", version, commit)
		return nil
	}

	if *serverFlag {
		opts := prombat.ServerOptions{
			Bind: *bindFlag,
			Port: *portFlag,
		}

		return prombat.RunServer(
			*namespaceFlag,
			prometheus.DefaultRegisterer.(*prometheus.Registry),
			opts,
		)
	}

	registry := prometheus.NewRegistry()
	err = registry.Register(prombat.NewCollector(*namespaceFlag))
	if err != nil {
		return err
	}

	gatherers := prometheus.Gatherers{registry}
	metricFamilies, err := gatherers.Gather()
	if err != nil {
		return err
	}

	var sb strings.Builder
	for _, mf := range metricFamilies {
		_, err := expfmt.MetricFamilyToText(&sb, mf)
		if err != nil {
			return err
		}
	}

	if *outputFlag != "" {
		return writeToFile(sb.String(), *outputFlag)
	}

	fmt.Print(sb.String())

	return nil
}

func setupSLog(levelStr string) error {
	var level slog.Level
	err := level.UnmarshalText([]byte(levelStr))
	if err != nil {
		return err
	}

	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
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
