package entity

import (
	"gorm.io/gorm"
)

type Customer struct{
	gorm.Model
	Name string `valid:"required~Name cannot be blank"`
	Email string
	CustomerID string `valid:"required~Name cannot be blank, matches(^[LMH]\\d{7}$)~CustomerID invalid format"`
}