package pswatch

import (
	"context"
	"errors"
	"os"
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

// StartProcess wraps the os.StartProcess function.
// Instead of returning the os.Process object, it returns a ProcessExit object, that gives the exit
// origin of this process when it's done. When the given ctx is done, it kills the process.
func StartProcess(ctx context.Context, name string, argv []string, attr *os.ProcAttr, pollMargin time.Duration) (ProcessExit, error) {
	ps, err := os.StartProcess(name, argv, attr)
	if err != nil {
		return nil, err
	}

	exit, err := WatchProcess(ctx, ps.Pid, pollMargin)
	if err != nil {
		return nil, errors.New("Process failed on startup")
	}

	return exit, nil
}
