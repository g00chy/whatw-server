package models

import (
	"time"
)

type Display struct {
	ID           uint
	Maker		string
	Model         string
	Size        uint8
	Hi          uint8
	Low          uint8
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Displays struct {
	data []Display
}

func (r *Displays) GetAll() []Display {
	return r.data
}