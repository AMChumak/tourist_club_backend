package model

import (
	"fmt"
	"github.com/jackc/pgx/v5/pgtype"
)

type Strain struct {
	Type     string
	Duration pgtype.Time
}

func (c *Strain) GetTimeAsString() string {
	us := c.Duration.Microseconds

	seconds := us / 1_000_000
	hours := seconds / 3600
	seconds %= 3600
	minutes := seconds / 60
	seconds %= 60

	return fmt.Sprintf("%02d:%02d", hours, minutes)
}
