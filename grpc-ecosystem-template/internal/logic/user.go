package logic

import (
	"context"
	"fmt"
	"grpc-ecosystem-template/api"
	"grpc-ecosystem-template/config"
	"grpc-ecosystem-template/internal/model"
	"grpc-ecosystem-template/utils"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/xiaomLee/go-plugin/db"
)

func CreateUser(ctx context.Context, user *api.User) error {
	bytes, err := hashPassword(user.Password)
	if err != nil {
		return fmt.Errorf("create user hash password err:%s", err.Error())
	}
	t, err := utils.ParseTime(user.Birthday)
	if err != nil {
		return fmt.Errorf("birthday format err:%s", err)
	}
	record := model.User{
		Id:         uint64(uuid.New().ID()),
		Name:       user.Name,
		Gender:     int32(user.Gender),
		Birthday:   utils.FormatTimeToDate(t),
		Status:     int32(api.User_STATUS_NORMAL),
		Password:   string(bytes),
		Account:    user.Account,
		Email:      user.Email,
		Tel:        user.Tel,
		ExtraInfo:  nil,
		CreateTime: time.Now().Unix(),
	}

	if err := db.GetDB("test").Create(&record).Error; err != nil {
		return fmt.Errorf("CreateUser insert db err:%s", err.Error())
	}
	user.Id = record.Id
	user.Password = ""
	user.Status = api.User_STATUS_NORMAL
	user.CreateTime = record.CreateTime
	return nil
}

// DeleteUser soft delete
func DeleteUser(ctx context.Context, id uint64) error {
	var user model.User
	db := db.GetDB("test").DB
	if config.Debug {
		db = db.Debug()
	}

	if err := db.Where("id=? and status<>?", id, int32(api.User_STATUS_DELETED)).First(&user).Error; err != nil {
		return fmt.Errorf("DeleteUser user not found, err:%s", err)
	}
	if err := db.Model(&user).Update("status", int32(api.User_STATUS_DELETED)).Error; err != nil {
		return fmt.Errorf("DeleteUser update user err:%s", err)
	}
	return nil
}

func GetUser(ctx context.Context, id uint64) (*api.User, error) {
	var user model.User
	db := db.GetDB("test").DB
	if config.Debug {
		db = db.Debug()
	}

	if err := db.Where("id=? and status<>?", id, int32(api.User_STATUS_DELETED)).First(&user).Error; err != nil {
		return nil, fmt.Errorf("GetUser user not found, err:%s", err)
	}
	return user.ConvertToProtoUser(), nil
}

func ListUser(ctx context.Context, params map[string]interface{}) ([]*api.User, error) {
	db := db.GetDB("test").DB
	if config.Debug {
		db = db.Debug()
	}
	db = db.Where("status<>?", int32(api.User_STATUS_DELETED))
	if id, ok := params["id"]; ok {
		db = db.Where("id=?", id)
	}
	if startTime, ok := params["start_time"]; ok {
		db = db.Where("create_time>=?", startTime)
	}
	if endTime, ok := params["end_time"]; ok {
		db = db.Where("create_time<?", endTime)
	}
	if name, ok := params["name"]; ok {
		db = db.Where("name like %?%", name)
	}
	if gender, ok := params["gender"]; ok {
		db = db.Where("gender in (?)", gender)
	}
	if status, ok := params["status"]; ok {
		db = db.Where("status in (?)", status)
	}
	if account, ok := params["account"]; ok {
		db = db.Where("account like %?%", account)
	}
	if email, ok := params["email"]; ok {
		db = db.Where("email like %?%", email)
	}
	if tel, ok := params["tel"]; ok {
		db = db.Where("tel=?", tel)
	}
	var records []*model.User
	if err := db.Find(&records).Error; err != nil {
		return nil, fmt.Errorf("list user db err:%s", err)
	}

	users := make([]*api.User, 0)
	for _, record := range records {
		users = append(users, record.ConvertToProtoUser())
	}
	return users, nil
}

func hashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func verifyPassword(password string, hashStr string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashStr), []byte(password)) == nil
}
