package model

type Admin struct {
	ID       string `json:"id" gorm:"primaryKey"`
	Login    string `json:"login" gorm:"not null;unique"`
	Password string `json:"password" gorm:"not null"`
}

type AdminOutput struct {
	ID    string `json:"id" gorm:"primaryKey"`
	Login string `json:"login" gorm:"not null;unique"`
}
