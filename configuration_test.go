package main

import (
	"runtime"
	"testing"

	"fmt"
	"io/ioutil"
	"os"

	"github.com/stretchr/testify/assert"
)

func TestConfigurationGetMaxNumberOfProcs(t *testing.T) {
	// Initialize
	c := &configuration{}
	numCPU := runtime.NumCPU()

	// Assert number of procs is always more than 0
	assert.Equal(t, 1, c.getMaxNumberOfProcs(0))

	// Assert number of procs is not bigger than number of CPUs
	assert.Equal(t, numCPU, c.getMaxNumberOfProcs(numCPU+1))
}

func TestConfigurationNew(t *testing.T) {
	// Create empty local conf
	f1, e := ioutil.TempFile(os.TempDir(), "bob_")
	defer os.Remove(f1.Name())
	assert.NoError(t, e)
	f1.WriteString("{}")

	// Get global configuration
	globalConf := newConfiguration(f1.Name())
	globalConf.parseGlobalConfiguration()

	// Create full local conf
	f2, e := ioutil.TempFile(os.TempDir(), "bob_")
	defer os.Remove(f2.Name())
	assert.NoError(t, e)
	sid := "test_stopwatch_id"
	f2.WriteString(fmt.Sprintf("{\"stopwatch\":{\"id\":\"%s\"}}", sid))

	// Create configuration with empty local conf
	c := newConfiguration(f2.Name())

	// Assert
	assert.Equal(t, globalConf.Logger.Prefix, c.Logger.Prefix)
	assert.NotEqual(t, globalConf.StopWatch.ID, c.StopWatch.ID)
	assert.Equal(t, sid, c.StopWatch.ID)
}
