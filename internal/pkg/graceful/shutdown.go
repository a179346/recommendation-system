package graceful

import (
	"os"
	"os/signal"
	"syscall"
)

func ShutDown() <-chan os.Signal {
	shutDownChannel := make(chan os.Signal, 1)
	signal.Notify(shutDownChannel, syscall.SIGINT, syscall.SIGTERM)
	return shutDownChannel
}
