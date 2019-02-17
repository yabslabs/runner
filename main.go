package main

import (
	"github.com/yabslabs/runner/common/logging"
)

func main() {
	logging.WithID("YABS-Runner-000").Info("runner started")
}
