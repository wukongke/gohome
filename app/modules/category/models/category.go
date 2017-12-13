package models

import (
	"gopkg.in/mgo.v2/bson"

	"work-codes/gohome/app/common"
	"work-codes/gohome/app/config"
)

var CategoryVO = common.DB(config.DBConfig.DbName).C("tb_category")

// Category 话题分类
type Category struct {
	Id        bson.ObjectId `bson:"_id"`
	Name      string        `bson:"name"`
	Sequence  int           `bson:"sequence"` // 同级别的分类可根据sequence的值来排序
	ParentID  string        `bson:"parentId"` // 直接父类ID
	Status    int           `bson:"status"`
	IsDeleted int           `bson:"isDeleted"`
	CreatedAt int           `bson:"createdAt"`
	UpdatedAt int           `bson:"updatedAt"`
	DeletedAt int           `bson:"deletedAt"`
}

// ToJSON 转成map
func (category *Category) ToJSON() map[string]interface{} {
	return map[string]interface{}{
		"id":        category.Id,
		"createdAt": category.CreatedAt,
		"updatedAt": category.UpdatedAt,
		"deletedAt": category.DeletedAt,
		"name":      category.Name,
		"sequence":  category.Sequence,
		"parentId":  category.ParentID,
		"status":    category.Status,
		"isDeleted": category.IsDeleted,
	}
}
