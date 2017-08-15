package pswatch

import (
	"context"
	"fmt"
	"os"
	"syscall"
	"time"
)

// DefaultPollMargin is the recommended time margin to poll against the system to know the status of a process
const DefaultPollMargin = time.Millisecond * 100

// WatchProcess returns a ProcessExit object, that gives the exit origin of this process when it's done.
// When the given ctx is done, it kills the process. If it didn't find the process, it returns an error.
func WatchProcess(ctx context.Context, pid int, pollMargin time.Duration) (ProcessExit, error) {
	p, err := os.FindProcess(pid)
	if err != nil {
		return nil, fmt.Errorf("No alive process with this PID")
	}

	exit := make(ProcessExit)

	go func() {
		defer close(exit)
		for {
			select {
			case <-ctx.Done():
				p.Kill()
				exit <- ProcessStoppedByUser
				return
			case <-time.After(pollMargin):
				err = p.Signal(syscall.Signal(0))
				if err != nil {
					exit <- ProcessHasDied
					return
				}
			}
		}
	}()

	return exit, nil
}
