package models

import (
	"errors"
	"time"
)

var _ = time.Thursday

type WpLitespeedImgOptm struct {
	Id             uint   `gorm:"column:id" form:"id" json:"id" comment:"" sql:"int(11) unsigned,PRI"`
	PostId         uint64 `gorm:"column:post_id" form:"post_id" json:"post_id" comment:"" sql:"bigint(20) unsigned,MUL"`
	OptmStatus     string `gorm:"column:optm_status" form:"optm_status" json:"optm_status" comment:"" sql:"varchar(64),MUL"`
	Src            string `gorm:"column:src" form:"src" json:"src" comment:"" sql:"varchar(1000)"`
	SrcpathMd5     string `gorm:"column:srcpath_md5" form:"srcpath_md5" json:"srcpath_md5" comment:"" sql:"varchar(128),MUL"`
	SrcMd5         string `gorm:"column:src_md5" form:"src_md5" json:"src_md5" comment:"" sql:"varchar(128),MUL"`
	Server         string `gorm:"column:server" form:"server" json:"server" comment:"" sql:"varchar(255)"`
	RootId         int    `gorm:"column:root_id" form:"root_id" json:"root_id" comment:"" sql:"int(11),MUL"`
	SrcFilesize    int    `gorm:"column:src_filesize" form:"src_filesize" json:"src_filesize" comment:"" sql:"int(11)"`
	TargetFilesize int    `gorm:"column:target_filesize" form:"target_filesize" json:"target_filesize" comment:"" sql:"int(11)"`
	TargetSaved    int    `gorm:"column:target_saved" form:"target_saved" json:"target_saved" comment:"" sql:"int(11)"`
	WebpFilesize   int    `gorm:"column:webp_filesize" form:"webp_filesize" json:"webp_filesize" comment:"" sql:"int(11)"`
	WebpSaved      int    `gorm:"column:webp_saved" form:"webp_saved" json:"webp_saved" comment:"" sql:"int(11)"`
}

func (m *WpLitespeedImgOptm) TableName() string {
	return "wp_litespeed_img_optm"
}

func (m *WpLitespeedImgOptm) One() (one *WpLitespeedImgOptm, err error) {
	one = &WpLitespeedImgOptm{}
	err = crudOne(m, one)
	return
}

func (m *WpLitespeedImgOptm) All(q *PaginationQuery) (list *[]WpLitespeedImgOptm, total uint, err error) {
	list = &[]WpLitespeedImgOptm{}
	total, err = crudAll(m, q, list)
	return
}

func (m *WpLitespeedImgOptm) Update() (err error) {
	where := WpLitespeedImgOptm{Id: m.Id}
	m.Id = 0
	return crudUpdate(m, where)
}

func (m *WpLitespeedImgOptm) Create() (err error) {
	m.Id = 0
	return mysqlDB.Create(m).Error
}

func (m *WpLitespeedImgOptm) Delete() (err error) {
	if m.Id == 0 {
		return errors.New("resource must not be zero value")
	}
	return crudDelete(m)
}
