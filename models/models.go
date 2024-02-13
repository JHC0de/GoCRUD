package models

import "gorm.io/gorm"

type Fact struct {
	gorm.Model
	Question string `json:"question" gorm:"text;not null;default:null"`
	Answer   string `json:"answer" gorm:"text;not null;default:null"`
}

type Wizard struct {
	gorm.Model
	Name   string `json:"name" gorm:"text;not null;default:null"`
	Type   string `json:"type" gorm:"text;not null;default:null"`
	Health int    `json:"health" gorm:"int;not null;default:null"`
	Armour int    `json:"armour" gorm:"int;not null;default:null"`
}
