package models

import (
	"errors"
	"time"
)

var _ = time.Thursday

type WpPost struct {
	Id                  uint       `gorm:"column:ID" form:"ID" json:"ID" comment:"" sql:"bigint(20) unsigned,PRI"`
	PostAuthor          uint64     `gorm:"column:post_author" form:"post_author" json:"post_author" comment:"" sql:"bigint(20) unsigned,MUL"`
	PostDate            *time.Time `gorm:"column:post_date" form:"post_date" json:"post_date,omitempty" comment:"" sql:"datetime"`
	PostDateGmt         *time.Time `gorm:"column:post_date_gmt" form:"post_date_gmt" json:"post_date_gmt,omitempty" comment:"" sql:"datetime"`
	PostContent         string     `gorm:"column:post_content" form:"post_content" json:"post_content" comment:"" sql:"longtext"`
	PostTitle           string     `gorm:"column:post_title" form:"post_title" json:"post_title" comment:"" sql:"text"`
	PostExcerpt         string     `gorm:"column:post_excerpt" form:"post_excerpt" json:"post_excerpt" comment:"" sql:"text"`
	PostStatus          string     `gorm:"column:post_status" form:"post_status" json:"post_status" comment:"" sql:"varchar(20)"`
	CommentStatus       string     `gorm:"column:comment_status" form:"comment_status" json:"comment_status" comment:"" sql:"varchar(20)"`
	PingStatus          string     `gorm:"column:ping_status" form:"ping_status" json:"ping_status" comment:"" sql:"varchar(20)"`
	PostPassword        string     `gorm:"column:post_password" form:"post_password" json:"post_password" comment:"" sql:"varchar(255)"`
	PostName            string     `gorm:"column:post_name" form:"post_name" json:"post_name" comment:"" sql:"varchar(200),MUL"`
	ToPing              string     `gorm:"column:to_ping" form:"to_ping" json:"to_ping" comment:"" sql:"text"`
	Pinged              string     `gorm:"column:pinged" form:"pinged" json:"pinged" comment:"" sql:"text"`
	PostModified        *time.Time `gorm:"column:post_modified" form:"post_modified" json:"post_modified,omitempty" comment:"" sql:"datetime"`
	PostModifiedGmt     *time.Time `gorm:"column:post_modified_gmt" form:"post_modified_gmt" json:"post_modified_gmt,omitempty" comment:"" sql:"datetime"`
	PostContentFiltered string     `gorm:"column:post_content_filtered" form:"post_content_filtered" json:"post_content_filtered" comment:"" sql:"longtext"`
	PostParent          uint64     `gorm:"column:post_parent" form:"post_parent" json:"post_parent" comment:"" sql:"bigint(20) unsigned,MUL"`
	Guid                string     `gorm:"column:guid" form:"guid" json:"guid" comment:"" sql:"varchar(255)"`
	MenuOrder           int        `gorm:"column:menu_order" form:"menu_order" json:"menu_order" comment:"" sql:"int(11)"`
	PostType            string     `gorm:"column:post_type" form:"post_type" json:"post_type" comment:"" sql:"varchar(20),MUL"`
	PostMimeType        string     `gorm:"column:post_mime_type" form:"post_mime_type" json:"post_mime_type" comment:"" sql:"varchar(100)"`
	CommentCount        int        `gorm:"column:comment_count" form:"comment_count" json:"comment_count" comment:"" sql:"bigint(20)"`
	ScrapyKey           string     `gorm:"column:scrapy_key" form:"scrapy_key" json:"scrapy_key" comment:"scrapy爬虫的key" sql:"varchar(255),MUL"`
}

func (m *WpPost) TableName() string {
	return "wp_posts"
}

func (m *WpPost) One() (one *WpPost, err error) {
	one = &WpPost{}
	err = crudOne(m, one)
	return
}

func (m *WpPost) All(q *PaginationQuery) (list *[]WpPost, total uint, err error) {
	list = &[]WpPost{}
	total, err = crudAll(m, q, list)
	return
}

func (m *WpPost) Update() (err error) {
	where := WpPost{Id: m.Id}
	m.Id = 0
	return crudUpdate(m, where)
}

func (m *WpPost) Create() (err error) {
	m.Id = 0
	return mysqlDB.Create(m).Error
}

func (m *WpPost) Delete() (err error) {
	if m.Id == 0 {
		return errors.New("resource must not be zero value")
	}
	return crudDelete(m)
}
