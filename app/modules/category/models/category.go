package models

import (
	"work-codes/gohome/app/common"
	"work-codes/gohome/app/config"
	"work-codes/gohome/app/libs"
)

var CategoryVO = common.DB(config.DBConfig.DbName).C("tb_category")

// Category 话题分类
type Category struct {
	libs.BaseModel
	Name     string `bson:"name"`
	Sequence int    `bson:"sequence"` // 同级别的分类可根据sequence的值来排序
	ParentID string `bson:"parentId"` // 直接父类ID
}

// ToJSON 转成map
func (model *Category) ToJSON() common.Map {
	return common.Map{
		"id":        model.Id,
		"name":      model.Name,
		"sequence":  model.Sequence,
		"parentId":  model.ParentID,
		"status":    model.Status,
		"isDeleted": model.IsDeleted,
		"createdAt": model.CreatedAt,
		"updatedAt": model.UpdatedAt,
		"deletedAt": model.DeletedAt,
	}
}
