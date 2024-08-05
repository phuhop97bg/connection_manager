package model

type Profile struct {
	ID        int    `json:"id" gorm:"column:id;primarykey;autoIncrement"`
	ProfileId string `json:"profileId" gorm:"column:profile_id"`
	FirstName string `json:"firstName" gorm:"column:first_name"`
	LastName  string `json:"lastName" gorm:"column:last_name"`
	Email     string `json:"email" gorm:"column:email"`
	Phone     string `json:"phone" gorm:"column:phone"`
}
