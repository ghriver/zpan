package model

import (
	"time"

	"gorm.io/gorm"
)

type Share struct {
	Id        int64      `json:"id"`
	Uid       int64      `json:"uid" gorm:"not null"`
	Alias     string     `json:"alias" gorm:"size:16;not null"`
	Matter    string     `json:"matter" gorm:"not null"`
	Name      string     `json:"name" gorm:"not null"`
	Type      string     `json:"type" gorm:"not null"`
	Secret    string     `json:"secret,omitempty" gorm:"size:16;not null"`
	Protected bool       `json:"protected" gorm:"-"`
	DownTimes int64      `json:"down_times" gorm:"not null"`
	ViewTimes int64      `json:"view_times" gorm:"not null"`
	ExpireAt  time.Time  `json:"expire_at" gorm:"not null"`
	CreateAt  time.Time  `json:"created" gorm:"autoCreateTime;not null"` // 这里的CreateAt应该是CreatedAt，但是先将错就错吧
	UpdateAt  time.Time  `json:"updated" gorm:"autoUpdateTime;not null"` // 这里的UpdateAt应该是UpdatedAt，但是先将错就错吧
	DeletedAt *time.Time `json:"-"`
}

func (Share) TableName() string {
	return "zp_share"
}

func (s *Share) AfterFind(*gorm.DB) (err error) {
	s.Protected = s.Secret != ""
	return
}
