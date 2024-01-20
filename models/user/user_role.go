package user

type URole struct {
	ID     uint `json:"id" gorm:"primaryKey" gorm:"comment:ID"`
	RoleID uint `json:"role_id" gorm:"comment:角色ID"`
	UserID uint `json:"user_id" gorm:"comment:用户ID"`
}
