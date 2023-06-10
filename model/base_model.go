package model

import "time"

type BaseModel struct {
	ID         int64     `json:"-" db:"id"`
	CreateTime time.Time `json:"create_time" db:"create_time"`
	UpdateTime time.Time `json:"update_time" db:"update_time"`
	DeleteTime time.Time `json:"-" db:"delete_time"`
	DeleteSate int       `json:"-" db:"delete_sate"`
}
