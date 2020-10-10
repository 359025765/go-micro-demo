package laracom_user_service

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

// 创建用户数据之前执行
func (model *User) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	return scope.SetColumn("Id", uuid.String())
}
