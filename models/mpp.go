package models

import "gorm.io/gorm"

type Mpp struct {
	gorm.Model
	EmployeeID    uint     `json:"employeeid"`
	Employee      Employee `json:"employee"`
	Period        string   `json:"period"`
	Numberreq     int      `json:"numberreq"`
	Status        int      `json:"status"`
	Reqheadcounts []Reqheadcount
}

type Reqheadcount struct {
	gorm.Model
	MppID            uint           `json:"mppid"`
	Mpp              Mpp            `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	EmployeeID       uint           `json:"employeeid"`
	Employee         Employee       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	LevelID          int            `json:"levelid"`
	Level            Level          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	GradeID          int            `json:"gradeid"`
	Grade            Grade          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Statusemployee   string         `json:"statusemployee"`
	Reasonhiring     string         `json:"reasonhiring"`
	Degree           string         `json:"degree"`
	Minexp           string         `json:"minexp"`
	JobDescriptionID uint           `json:"jobdescriptionid"`
	JobDescription   JobDescription `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Specification    string         `json:"specification"`
	Gender           string         `json:"gender"`
	Age              string         `json:"age"`
	Maritalstatus    string         `json:"maritalstatus"`
	Recruitmenttype  string         `json:"recruitmenttype"`
	Status           string         `json:"status"`
}
