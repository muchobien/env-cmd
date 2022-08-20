package build

import (
	"runtime/debug"
	"time"
)

// Version is dynamically set by the toolchain or overridden by the Makefile.
var Version = "DEV"

// Date is dynamically set at build time in the Makefile.
var Date = "" // YYYY-MM-DD

// Compiled is dynamically set at build time in the Makefile.
var Compiled time.Time // YYYY-MM-DD

func init() {
	if Version == "DEV" {
		if info, ok := debug.ReadBuildInfo(); ok && info.Main.Version != "(devel)" {
			Version = info.Main.Version
		}
	}

	if Date != "" {
		if parsed, err := time.Parse("2006-01-02", Date); err == nil {
			Compiled = parsed
		}
	} else {
		Compiled = time.Now()
	}
}
