package models

import (
	"work-codes/gohome/app/common"
	"work-codes/gohome/app/config"
	"work-codes/gohome/app/libs"
)

var UserVO = common.DB(config.DBConfig.DbName).C("tb_user")

// 用户表
type User struct {
	libs.BaseModel
	Name string `bson:"name"`
}

// ToJSON 转成map
func (model *User) ToJSON() common.Map {
	return common.Map{
		"id":        model.Id,
		"name":      model.Name,
		"status":    model.Status,
		"isDeleted": model.IsDeleted,
		"createdAt": model.CreatedAt,
		"updatedAt": model.UpdatedAt,
		"deletedAt": model.DeletedAt,
	}
}
