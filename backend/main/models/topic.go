package models

import (
	"database/sql"
	"time"
)

type Topic struct {
	ID          uint64       `db:"id"`
	Title		string		 `db:"title"`
	Description string       `db:"description"`
	ImgUrl	    string	 	 `db:"imag_url"`
	FirstCategory string     `db:"first_category"`
	secondCategory string    `db:"second_category"`
	Articles 	[]Article    `db:"articles"`
	CreatedAt   time.Time    `db:"created_at"`
	UpdatedAt   sql.NullTime `db:"updated_at"`
	DeletedAt   sql.NullTime `db:"deleted_at"`	
}