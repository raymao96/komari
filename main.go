package main

import (
	"log/slog"
	"os"

	"github.com/raymao96/komari/cmd"
	"github.com/raymao96/komari/utils"
	logger "github.com/raymao96/komari/utils/log"
)

func main() {
	// Version probes are consumed by the updater and must contain no log prefix.
	if len(os.Args) > 1 && os.Args[1] == "version" {
		cmd.Execute()
		return
	}
	if utils.VersionHash == "unknown" {
		logger.Setup(slog.LevelDebug)
	} else {
		logger.Setup(slog.LevelInfo)
	}

	logger.Infof("server", "Lite %s (%s)", utils.CurrentVersion, utils.VersionHash)

	cmd.Execute()
}
