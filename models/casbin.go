package models

import "gorm.io/gorm"

type Casbin struct {
	gorm.Model
	PType string `gorm:"type:varchar(100);uniqueIndex:casbin_rule_key"`
	V0    string `gorm:"type:varchar(100);uniqueIndex:casbin_rule_key"`
	V1    string `gorm:"type:varchar(100);uniqueIndex:casbin_rule_key"`
	V2    string `gorm:"type:varchar(100);uniqueIndex:casbin_rule_key"`
	V3    string `gorm:"type:varchar(100);uniqueIndex:casbin_rule_key"`
	V4    string `gorm:"type:varchar(100);uniqueIndex:casbin_rule_key"`
	V5    string `gorm:"type:varchar(100);uniqueIndex:casbin_rule_key"`
}
