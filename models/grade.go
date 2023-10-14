package models

import "gorm.io/gorm"

type Grade struct {
	gorm.Model
	Grade         int        `gorm:"unique" json:"grade"`
	Min           int64      `json:"min"`
	Max           int64      `json:"max"`
	Struktur      string     `json:"struktur"`
	Employees     []Employee `gorm:"foreignKey:GradeID;references:Grade"`
	Reqheadcounts []Reqheadcount
}

type Increament struct {
	gorm.Model
	Percent string `json:"percent"`
}
