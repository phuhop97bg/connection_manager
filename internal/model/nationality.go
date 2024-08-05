package model

type Nationality struct {
	ID   int    `json:"id" gorm:"column:id;primarykey;autoIncrement"`
	Name string `json:"name" gorm:"column:name;"`
	Code string `json:"code" gỏm:"code"`
}

type NationalityTabler interface {
	TableName() string
}

func (Nationality) TableName() string {
	return "nationalities"
}
