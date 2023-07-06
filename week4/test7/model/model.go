package model

import "gorm.io/gorm"

type Employer struct {
	gorm.Model
	Name      string
	CompanyID string
	Campany   Company
}

type CreditCard struct {
	gorm.Model
	Number     string
	EmployerID uint
}

type Company struct {
	gorm.Model
	Name string
}
