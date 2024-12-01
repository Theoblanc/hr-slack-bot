package commands

import "time"

type ClockIn struct {
	EmployeeID uint
	Location   string
	Timestamp  time.Time
}

type ClockOut struct {
	EmployeeID uint
	Location   string
	Timestamp  time.Time
}
