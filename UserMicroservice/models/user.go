package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Nombre    string `json:"nombre" gorm:"not null"`
	Email     string `json:"email" gorm:"not null;unique"`
	Password  string `json:"password" gorm:"not null"`
	Telefono  string `json:"telefono" gorm:"not null;unique"`
	Pais      string `json:"pais,omitempty"`
	Ciudad    string `json:"ciudad,omitempty"`
	Direccion string `json:"direccion,omitempty"`
}
