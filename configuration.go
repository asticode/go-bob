package main

import (
	"bytes"
	"encoding/json"
	"runtime"

	"math"

	"github.com/asticode/go-logger/logger"
	"github.com/asticode/go-stopwatch/stopwatch"
	extendederror "github.com/asticode/go-toolbox/error"
	extendedfile "github.com/asticode/go-toolbox/file"
)

// Configuration represents a JSON-friendly structure of the configuration attributes
type configuration struct {
	Bob       configurationBob        `json:"bob"`
	Logger    logger.Configuration    `json:"logger"`
	StopWatch stopwatch.Configuration `json:"stopwatch"`
}

// NewConfiguration parses and merges global and local configurations, and returns a pointer to a Configuration object
func newConfiguration(cp string) *configuration {
	// Initialize
	c := &configuration{}

	// Parse global configuration
	c.parseGlobalConfiguration()

	// Parse local configuration
	c.parseLocalConfiguration(cp)

	// Update max number of procs
	runtime.GOMAXPROCS(c.getMaxNumberOfProcs(c.Bob.MaxNumberOfProcs))

	// Return
	return c
}

// ParseGlobalConfiguration parses the global configuration
func (c *configuration) parseGlobalConfiguration() {
	gc, err := Asset("resources/config/global.json")
	extendederror.ProcessError(err)
	err = json.NewDecoder(bytes.NewReader(gc)).Decode(c)
	extendederror.ProcessError(err)
}

func (c *configuration) parseLocalConfiguration(cp string) {
	err := json.NewDecoder(bytes.NewReader(extendedfile.GetContentBytes(cp))).Decode(c)
	extendederror.ProcessError(err)
}

func (c *configuration) getMaxNumberOfProcs(confValue int) int {
	mp := int(math.Max(float64(confValue), float64(1)))
	return int(math.Min(float64(mp), float64(runtime.NumCPU())))
}

type configurationBob struct {
	MaxNumberOfProcs int `json:"max_number_of_procs"`
}
