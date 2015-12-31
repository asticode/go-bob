package main

import (
	"flag"
	"fmt"

	"github.com/asticode/go-logger/logger"
	"github.com/asticode/go-stopwatch/stopwatch"
	extendederror "github.com/asticode/go-toolbox/error"
	"github.com/asticode/go-virtualkeyboard/virtualkeyboard"
	"time"
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
	sw := stopwatch.NewStopwatchFromConfiguration(c.StopWatch).SetIsEnabled(true)
	sw.AddEvent("Init", "Stopwatch has been created")

	//>> TESTS BEGIN HERE <<\\

	// Create virtual keyboard
	vk := virtualkeyboard.NewVirtualKeyboard().SetDelayBetweenPresses(time.Duration(50) * time.Millisecond)

	// Sleep
	time.Sleep(time.Duration(4) * time.Second)

	// Press
	fmt.Println(vk.Write("Salut Quentin, tu vas bien ?"))

	//>> TESTS END HERE <<\\

	// Log stopwatch
	l.Info(sw.String())
}
