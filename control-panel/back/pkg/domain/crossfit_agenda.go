package domain

import (
	"time"
)

const (
	Failed ProcessStatus = iota + 1
	Finished
	Working
)

type (
	ProcessStatus int

	Status struct {
		ID          ProcessStatus
		Date        time.Time
		Description string
		Complete    bool
	}

	Credentials struct {
		URI   string
		Token string
	}
)
