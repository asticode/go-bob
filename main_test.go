package main

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/asticode/go-logger/logger"
	extendederror "github.com/asticode/go-toolbox/error"
	"github.com/stretchr/testify/assert"
)

func testCatchPreparePanic() {
	defer extendederror.Catch(catchPrepare, map[string]interface{}{})
	panic("error")
}

// TestCatchPrepare checks that error are properly handled and logged in stdout when executing "prepare"
func TestCatchPrepare(t *testing.T) {
	// Create file
	file, e := ioutil.TempFile(os.TempDir(), "bob_")
	defer os.Remove(file.Name())
	assert.NoError(t, e)

	// Save and reroute stdout
	old := os.Stdout
	os.Stdout = file

	// Panic
	testCatchPreparePanic()

	// Assert
	c, e := ioutil.ReadFile(file.Name())
	assert.NoError(t, e)
	assert.Equal(t, "CRITICAL: error\n", string(c))

	// Restore stdout
	os.Stdout = old
}

func testCatchRunPanic(l logger.Logger) {
	defer extendederror.Catch(catchRun, map[string]interface{}{
		"logger": l,
	})
	panic("error")
}

// TestCatchCreateServer checks that error are properly handled and logged when executing "createServer"
func TestCatchCreateServer(t *testing.T) {
	// Create file
	file, e := ioutil.TempFile(os.TempDir(), "bob_")
	defer os.Remove(file.Name())

	// Create logger
	l, e := logger.NewLogger("prefix", logger.LevelError)
	assert.NoError(t, e)
	l.SetFormat("[%level%]%message%")

	// Add file handler
	h, e := logger.NewHandlerFile(file.Name())
	assert.NoError(t, e)
	l.AddHandler("file", h)

	// Panic
	testCatchRunPanic(l)

	// Assert
	c, e := ioutil.ReadFile(file.Name())
	assert.NoError(t, e)
	assert.Equal(t, "[CRITICAL]error\n", string(c))
}
