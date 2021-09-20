package users

import "gorm.io/gorm"

const (
	AdminRole Role = iota
	InvestorRole
	CompanyRole
)

type Role int

type User struct {
	gorm.Model `json:"-"`
	Name       string `json:"name"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	SysRole    Role   `json:"role"`
}
