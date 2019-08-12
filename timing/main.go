package timing

import (
	"time"
	"os"
	"os/signal"
	"syscall"
)

func SetInterval(f func(), ms int, async bool) chan bool {
	interval := time.Duration(ms) * time.Millisecond
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	clear := make(chan bool)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	defer signal.Stop(sig)

	go func() {
		for {

			select {
			case <-ticker.C:
				if async {
					go f()
				} else {
					f()
				}
			case <- clear:
				ticker.Stop()
				return

			case s := <-sig:
				switch s {
				// WindowsだとSIGKILL以外効かない
				case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
					return
				}
			}

		}
	}()
	return clear
}