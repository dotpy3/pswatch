package pswatch

import (
	"context"
	"errors"
	"fmt"
	"os"
	"syscall"
	"time"
)

// ProcessExit is a channel type, sending the exit code of a process when it's complete.
// In a ProcessExit is first piped the exit code of a process - the channel is then closed.
type ProcessExit chan int

const (
	// ProcessStoppedByUser is the value sent through a ProcessExit when the process was stopped by
	// the manager, following a signal.
	ProcessStoppedByUser int = 1
	// ProcessHasDied is the value sent through a ProcessExit when the process wasn't stopped by user
	// signal, but for another reason: stopped by the system, stopped on its own...
	ProcessHasDied int = 2
)

// DefaultPollMargin is the recommended time margin to poll against the system to know the status of a process
const DefaultPollMargin = time.Millisecond * 100

// WatchProcess returns a ProcessExit object, that gives the exit origin of this process when it's done.
// When the given ctx is done, it kills the process. If it didn't find the process, it returns an error.
func WatchProcess(ctx context.Context, pid int, pollMargin time.Duration) (ProcessExit, error) {
	p, err := os.FindProcess(pid)
	if err != nil {
		return nil, err
	}
	err = p.Signal(syscall.Signal(0))
	if err != nil {
		return nil, errors.New("No process found with that PID")
	}

	exit := make(ProcessExit)
	fmt.Println("hello")
	go func() {
		defer close(exit)
		fmt.Println("hi")
		for {
			time.Sleep(pollMargin)
			select {
			case <-ctx.Done():
				p.Kill()
				exit <- ProcessStoppedByUser
				return
			default:
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
