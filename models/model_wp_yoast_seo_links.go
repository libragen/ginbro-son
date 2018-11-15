package models

import (
	"errors"
	"time"
)

var _ = time.Thursday

type WpYoastSeoLink struct {
	Id           uint   `gorm:"column:id" form:"id" json:"id" comment:"" sql:"bigint(20) unsigned,PRI"`
	Url          string `gorm:"column:url" form:"url" json:"url" comment:"" sql:"varchar(255)"`
	PostId       uint64 `gorm:"column:post_id" form:"post_id" json:"post_id" comment:"" sql:"bigint(20) unsigned,MUL"`
	TargetPostId uint64 `gorm:"column:target_post_id" form:"target_post_id" json:"target_post_id" comment:"" sql:"bigint(20) unsigned"`
	Type         string `gorm:"column:type" form:"type" json:"type" comment:"" sql:"varchar(8)"`
}

func (m *WpYoastSeoLink) TableName() string {
	return "wp_yoast_seo_links"
}

func (m *WpYoastSeoLink) One() (one *WpYoastSeoLink, err error) {
	one = &WpYoastSeoLink{}
	err = crudOne(m, one)
	return
}

func (m *WpYoastSeoLink) All(q *PaginationQuery) (list *[]WpYoastSeoLink, total uint, err error) {
	list = &[]WpYoastSeoLink{}
	total, err = crudAll(m, q, list)
	return
}

func (m *WpYoastSeoLink) Update() (err error) {
	where := WpYoastSeoLink{Id: m.Id}
	m.Id = 0
	return crudUpdate(m, where)
}

func (m *WpYoastSeoLink) Create() (err error) {
	m.Id = 0
	return mysqlDB.Create(m).Error
}

func (m *WpYoastSeoLink) Delete() (err error) {
	if m.Id == 0 {
		return errors.New("resource must not be zero value")
	}
	return crudDelete(m)
}
