package pswatch

import (
	"context"
	"testing"
)

func TestUnknownPid(t *testing.T) {
	_, err := WatchProcess(context.Background(), -1, DefaultPollMargin)
	if err == nil {
		t.Error("WatchProcess didn't fail with a PID of -1")
	}
}
