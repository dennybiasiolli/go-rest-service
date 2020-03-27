package models

import "time"

type ResUser struct {
	ID            int       `gorm:"PRIMARY_KEY;NOT NULL;AUTO_INCREMENT"`
	Active        bool      `gorm:"DEFAULT:true"`
	Login         string    `gorm:"type:varchar;NOT NULL"`
	Password      string    `gorm:"type:varchar;DEFAULT:NULL"`
	CompanyID     int       `gorm:"NOT NULL"`
	PartnerID     int       `gorm:"NOT NULL"`
	Signature     string    `gorm:"Type:text"`
	ActionID      int       `gorm:"default:null"`
	Share         bool      ``
	CreateUID     int       ``
	CreateDate    time.Time `gorm:"Type:timestamp without time zone"`
	WriteUID      int       ``
	WriteDate     time.Time `gorm:"Type:timestamp without time zone"`
	PasswordCrypt string    `gorm:"Type:varchar"`
}

// // set table name
// func (ResUser) TableName() string {
// 	return "res_users"
// }
