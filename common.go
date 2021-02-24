package common

import (
	"fmt"
	"time"
)

type (
	ProcAttributes struct {
		Name      string
		Cmdline   []string
		Username  string
		PID       int
		PPID      int
		StartTime time.Time
	}

	MatchNamer interface {
		// MatchAndName returns false if the match failed, otherwise
		// true and the resulting name.
		MatchAndName(ProcAttributes) (bool, string)
		fmt.Stringer
	}
)
