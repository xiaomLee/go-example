package model

import (
	"database/sql/driver"
	jsoniter "github.com/json-iterator/go"
	"time"
)

type User struct {
	Id       string    `gorm:"primaryKey;comment:记录ID" json:"id"`
	Name     string    `gorm:"comment:名称" json:"name"`
	Gender   uint32    `gorm:"comment:设备名称" json:"gender"`
	Birthday time.Time `gorm:"comment:Birthday" json:"birthday"`
	Status   uint32    `gorm:"comment:状态" json:"status"`

	Password  string `gorm:"comment:pwd" json:"password"`
	Account   string `gorm:"comment:account, 秒" json:"account"`
	Email     string `gorm:"comment:email" json:"email"`
	Tel       string `gorm:"comment:tel" json:"tel"`
	ExtraInfo Extra  `gorm:"type:json;comment:扩展信息" json:"extra_info"`

	CreateTime int64 `gorm:"autoCreateTime;comment:创建时间" json:"create_time"`
	UpdateTime int64 `gorm:"autoUpdateTime;comment:更新时间" json:"update_time"`
	UpdateUser int64 `gorm:"comment:更新人" json:"update_user"`
}

type Extra map[string]string

func (t Extra) Value() (driver.Value, error) {
	return jsoniter.Marshal(t)
}

func (t *Extra) Scan(v interface{}) (err error) {
	return jsoniter.Unmarshal(v.([]uint8), t)
}
