package models

import (
	"time"

	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	EmployeeId     string     `json:"employeeid" gorm:"unique"`
	Name           string     `json:"name"`
	Email          string     `json:"email" gorm:"unique"`
	GradeId        int        `json:"gradeid"`
	Grade          Grade      `gorm:"references:GradeId"`
	Salary         int64      `json:"salary"`
	Statusemployee string     `json:"statusemployee"`
	Joindate       *time.Time `json:"joindate"`
	Resigndate     *time.Time `json:"resigndate"`
	Endcontract    *time.Time `json:"endcontract"`
	Address        string     `json:"address"`
	Ciaddress      string     `json:"ciaddress"`
	Norek          string     `json:"norek"`
	Noktp          int        `json:"noktp"`
	Npwp           string     `json:"npwp"`
	Kis            string     `json:"kis"`
	Kpj            string     `json:"kpj"`
	Ptkp           string     `json:"ptkp"`
	Phone          string     `json:"phone"`
	Birthplace     string     `json:"birthplace"`
	Birthdate      *time.Time `json:"birthdate"`
	Gender         string     `json:"gender"`
	Religion       string     `json:"religion"`
	Marital        string     `json:"marital"`
	National       string     `json:"national"`
}

type Candidate struct {
	gorm.Model
}