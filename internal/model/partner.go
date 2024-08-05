package model

type Partner struct {
	Id          int    `json:"id" db:"id" gorm:"column:id;primarykey;autoIncrement"`
	PartnerId   string `json:"partnerId" db:"partner_id" gorm:"column:partner_id"`
	PartnerName string `json:"partnerName" db:"partner_name" gorm:"column:partner_name"`
}
