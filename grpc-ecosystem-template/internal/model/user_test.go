package model

import (
	"testing"
	"time"

	"github.com/xiaomLee/go-plugin/db"
)

func InitDB(t *testing.T) {
	if err := db.AddDB(db.SQLITE, "test", "file:../../test.db?_auth&_auth_user=admin&_auth_pass=admin&_auth_crypt=sha1"); err != nil {
		t.Fatal(err)
	}
}
func TestCreateUser(t *testing.T) {
	InitDB(t)
	user := User{
		Id:         "2",
		Name:       "two",
		Gender:     2,
		Birthday:   time.Now(),
		Status:     1,
		Password:   "admin",
		Account:    "admin",
		Email:      "admin@gmail.com",
		Tel:        "12300001111",
		ExtraInfo:  map[string]string{"a": "a", "b": "b"},
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
		UpdateUser: 0,
	}
	if err := db.MustGetDB("test").Create(&user).Error; err != nil {
		t.Error()
	}
}
