package models

import "github.com/dgrijalva/jwt-go"

type User struct {
	ID          int    `json:"id"`
	Apartment   int    `json:"apartment,omitempty"`
	Email       string `json:"email,omitempty"`
	RawPassword string `json:"password,omitempty" gorm:"-"`
	Password    string `json:"-"`
	Phone       string `json:"phone,omitempty"`
	FirstName   string `json:"first_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	Role        string `json:"role,omitempty"`
	ResidenceID int    `json:"residence_id,omitempty"`
	BuildingID  int    `json:"building_id,omitempty"`
	Token       string `json:"token" gorm:"-"`
}

func (User) TableName() string {
	return "user"
}

type Token struct {
	ID int
	jwt.StandardClaims
}

//type Residence struct {
//	ID      int    `json:"id"`
//	Name    string `json:"name"`
//	Address string `json:"address"`
//}
//
//type Building struct {
//	ID          int `json:"id"`
//	ResidenceID int `json:"residence_id"`
//	Name        int `json:"name"`
//	Address     int `json:"address"`
//}
