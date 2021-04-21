package model



//用户
type UsrUser struct {
    //
    Id string  `gorm:"column:id"`
    //
    Name string  `gorm:"column:name"`
    //
    Password string  `gorm:"column:password"`
    
}

//用户表名
func (UsrUser) TableName() string {
	return "usr_user"
}
