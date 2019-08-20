package models

type User struct {
	ID          int    `json:"id"`
	Appartment  int    `json:"appartment,omitempty"`
	Email       string `json:"email,omitempty"`
	Password    string `json:"password,omitempty"`
	Phone       string `json:"phone,omitempty"`
	FirstName   string `json:"first_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	Role        string `json:"role,omitempty"`
	ResidenceID int    `json:"residence_id,omitempty"`
	BuildingID  int    `json:"building_id,omitempty"`
}

func (User) TableName() string {
	return "user"
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
