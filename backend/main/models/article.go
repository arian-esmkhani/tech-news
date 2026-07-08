package models

import (
	"database/sql"
	"time"
)

type Article struct {
	ID            uint64       `db:"id"`
	Head      	  string       `db:"head"`
	Body	      string       `db:"body"`
	ImgUrl        string       `db:"imag_url"`
	Sec        	  int64        `db:"sec"` //order
	Topic		  Topic        `db:"topic_id"`
	CreatedAt     time.Time    `db:"created_at"`
	UpdatedAt     sql.NullTime `db:"updated_at"`
	DeletedAt     sql.NullTime `db:"deleted_at"`
}