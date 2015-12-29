package main

import (
	"flag"
	"fmt"

	"github.com/asticode/go-logger/logger"
	extendederror "github.com/asticode/go-toolbox/error"
	"github.com/asticode/go-stopwatch/stopwatch"
)

func main() {
	// Defer
	defer extendederror.Catch(catchPrepare, map[string]interface{}{})

	// Parse flags.
	configPath := flag.String("config", "", "Path to the local configuration file")
	flag.Parse()

	// Run
	prepare(*configPath)
}

func catchPrepare(err interface{}, args map[string]interface{}) {
	fmt.Println("CRITICAL: " + extendederror.ParseError(err))
}

func prepare(configPath string) {
	// Create configuration
	c := newConfiguration(configPath)

	// Create logger
	l, err := logger.NewLoggerFromConfiguration(c.Logger)
	extendederror.ProcessError(err)

	// Init.
	defer extendederror.Catch(catchRun, map[string]interface{}{
		"logger": l,
	})
	run(*c, l)
}

func catchRun(err interface{}, args map[string]interface{}) {
	// Initialize
	l := args["logger"].(logger.Logger)

	// Log message
	l.Critical(extendederror.ParseError(err))
}

func run(c configuration, l logger.Logger) {
	// Create stopwatch
	s := stopwatch.NewStopwatchFromConfiguration(c.StopWatch).SetIsEnabled(true)
	s.AddEvent("Init", "Stopwatch has been created")

	l.Info(s.String())
}
