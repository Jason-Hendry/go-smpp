package signal

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func Wait() chan bool {
	signals := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-signals
		fmt.Println("Signal received: ", sig)
		done <- true
	}()
	return done
}
