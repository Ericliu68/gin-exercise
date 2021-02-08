package models

// mysql对应数据库
type Users struct {
	Id       int    `gorm:"column:id;primaryKey"`
	UserName string `gorm:"column:name;unique"`
	Password string `gorm:"column:password"`
	RoleID   int    `gorm:"column:role"`
	Role     string `gorm:"column:role"`
	Avatar   []byte `gorm:"column:avatar"`
}

// 登录form data 绑定结构体
type UserLogin struct {
	UserName string `form:"user" json:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

type Role struct {
	Id         int    `gorm:"column:id;primaryKey"`
	Role       string `gorm:"column:role;unique"`
	Permission int    `gorm:"column:permission"`
}
