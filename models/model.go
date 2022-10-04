package models

type Iphone struct {
	Id      int    `json:"id" gorm:"primaryKey"`
	Name    string `json:"name"`
	Model   string `json:"model"`
	Feature string `json:"feature"`
}
