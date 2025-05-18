package model

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Championship struct {
	Id    int32
	Title string
	Date  pgtype.Date
}

func (c *Championship) GetDateAsString() string {
	t := c.Date.Time
	return t.Format("2006-01-02") // Формат YYYY-MM-DD
}
