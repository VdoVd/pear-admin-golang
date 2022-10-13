package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Role struct {
	ID        uint       `json:"id"`
	RoleName  string     `json:"role_name"`
	Detail    string     `json:"detail"`
	CreateId  int        `json:"create_id"`
	Status    int        `json:"status"`
	UpdateId  int        `json:"update_id"`
	RoleAuths []RoleAuth `json:"-" gorm:"foreignkey:RoleID"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

func (r *Role) BeforeCreate(scope *gorm.Scope) {
	scope.SetColumn("Status", 1)
}

func (r *Role) BeforeUpdate(scope *gorm.Scope) {

}

type RoleEditShow struct {
	ID       int
	RoleName string
	Status   int
	Checked  int
}
