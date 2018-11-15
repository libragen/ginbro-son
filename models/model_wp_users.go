package models

import (
	"errors"
	"time"
)

var _ = time.Thursday

type WpUser struct {
	Id                uint       `gorm:"column:ID" form:"ID" json:"ID" comment:"" sql:"bigint(20) unsigned,PRI"`
	UserLogin         string     `gorm:"column:user_login" form:"user_login" json:"user_login" comment:"" sql:"varchar(60),MUL"`
	UserPass          string     `gorm:"column:user_pass" form:"user_pass" json:"user_pass" comment:"" sql:"varchar(255)"`
	UserNicename      string     `gorm:"column:user_nicename" form:"user_nicename" json:"user_nicename" comment:"" sql:"varchar(50),MUL"`
	UserEmail         string     `gorm:"column:user_email" form:"user_email" json:"user_email" comment:"" sql:"varchar(100),MUL"`
	UserUrl           string     `gorm:"column:user_url" form:"user_url" json:"user_url" comment:"" sql:"varchar(100)"`
	UserRegistered    *time.Time `gorm:"column:user_registered" form:"user_registered" json:"user_registered,omitempty" comment:"" sql:"datetime"`
	UserActivationKey string     `gorm:"column:user_activation_key" form:"user_activation_key" json:"user_activation_key" comment:"" sql:"varchar(255)"`
	UserStatus        int        `gorm:"column:user_status" form:"user_status" json:"user_status" comment:"" sql:"int(11)"`
	DisplayName       string     `gorm:"column:display_name" form:"display_name" json:"display_name" comment:"" sql:"varchar(250)"`
}

func (m *WpUser) TableName() string {
	return "wp_users"
}

func (m *WpUser) One() (one *WpUser, err error) {
	one = &WpUser{}
	err = crudOne(m, one)
	return
}

func (m *WpUser) All(q *PaginationQuery) (list *[]WpUser, total uint, err error) {
	list = &[]WpUser{}
	total, err = crudAll(m, q, list)
	return
}

func (m *WpUser) Update() (err error) {
	where := WpUser{Id: m.Id}
	m.Id = 0
	return crudUpdate(m, where)
}

func (m *WpUser) Create() (err error) {
	m.Id = 0
	return mysqlDB.Create(m).Error
}

func (m *WpUser) Delete() (err error) {
	if m.Id == 0 {
		return errors.New("resource must not be zero value")
	}
	return crudDelete(m)
}
