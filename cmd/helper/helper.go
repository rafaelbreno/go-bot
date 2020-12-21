package helper

import (
	"net"
	"os"
	"sync"
)

func WaitForCtrlC() {
	var wait sync.WaitGroup

	wait.Add(1)

	var wait_channel chan os.Signal

	wait_channel = make(chan os.Signal, 1)

	go func() {
		<-wait_channel
		wait.Done()
	}()

	wait.Wait()
}

func WriteTwitch(msg string, conn net.Conn) {
	msg += "\r\n"
	conn.Write([]byte(msg))
}
