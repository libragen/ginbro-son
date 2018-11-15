package models

import (
	"errors"
	"time"
)

var _ = time.Thursday

type WpLitespeedOptimizer struct {
	Id       uint   `gorm:"column:id" form:"id" json:"id" comment:"" sql:"int(11),PRI"`
	HashName string `gorm:"column:hash_name" form:"hash_name" json:"hash_name" comment:"hash.filetype" sql:"varchar(60),UNI"`
	Src      string `gorm:"column:src" form:"src" json:"src" comment:"full url array set" sql:"text"`
	Dateline int    `gorm:"column:dateline" form:"dateline" json:"dateline" comment:"" sql:"int(11),MUL"`
	Refer    string `gorm:"column:refer" form:"refer" json:"refer" comment:"The container page url" sql:"varchar(255)"`
}

func (m *WpLitespeedOptimizer) TableName() string {
	return "wp_litespeed_optimizer"
}

func (m *WpLitespeedOptimizer) One() (one *WpLitespeedOptimizer, err error) {
	one = &WpLitespeedOptimizer{}
	err = crudOne(m, one)
	return
}

func (m *WpLitespeedOptimizer) All(q *PaginationQuery) (list *[]WpLitespeedOptimizer, total uint, err error) {
	list = &[]WpLitespeedOptimizer{}
	total, err = crudAll(m, q, list)
	return
}

func (m *WpLitespeedOptimizer) Update() (err error) {
	where := WpLitespeedOptimizer{Id: m.Id}
	m.Id = 0
	return crudUpdate(m, where)
}

func (m *WpLitespeedOptimizer) Create() (err error) {
	m.Id = 0
	return mysqlDB.Create(m).Error
}

func (m *WpLitespeedOptimizer) Delete() (err error) {
	if m.Id == 0 {
		return errors.New("resource must not be zero value")
	}
	return crudDelete(m)
}
