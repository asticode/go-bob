package main

import (
	"flag"
	"fmt"

	"github.com/asticode/go-logger/logger"
	"github.com/asticode/go-stopwatch/stopwatch"
	extendederror "github.com/asticode/go-toolbox/error"
	_ "github.com/asticode/go-keyboardemulator/keyboardemulator"
	_ "os/exec"
	"github.com/asticode/go-texttospeech/texttospeech"
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

	/*
	// Create keyboard emulator
	vk := keyboardemulator.NewKeyboardEmulator().SetAddRealLifeDelayBetweenPresses(true)

	// Switch focus and select all content
	vk.Press(
		// Switch focus
		keyboardemulator.Keys{
			keyboardemulator.KeyAlt,
			keyboardemulator.KeyTab,
		},
		// Select all
		keyboardemulator.Keys{
			keyboardemulator.KeyControl,
			keyboardemulator.KeyA,
		},
	)
	sw.AddEvent("First press", "")

	// Write
	vk.Write("Bonjour Quentin, comment vas-tu ?")
	sw.AddEvent("Second press", "")

	// Open Explorer
	exec.Command("cmd", "/C explorer .").Output()
	*/

	// Create text to speech
	tts := texttospeech.NewTextToSpeech()

	// Say
	tts.Say("Good morning sir, how may I help you today?")

	//>> TESTS END HERE <<\\

	// Log stopwatch
	l.Info("Done")
}
