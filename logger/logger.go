package logger

import (
	"io"
	"log"
	"os"
)

var (
	// Info logger, by deault print to os.Stdout
	Info = log.New(os.Stdout, "INFO\t", log.LstdFlags)
	// Debug logger, must be enabled by flag
	Debug = log.New(io.Discard, "DEBUG\t", log.LstdFlags)
)

// Init is initializing the Debug logger
func Init(verbose bool) {
	if verbose {
		Debug = log.New(os.Stdout, "DEBUG\t", log.LstdFlags)
	}
}
