package model



//产品
type IhrGruleObject struct {
    //
    Id int  `gorm:"column:id"`
    //
    Code string  `gorm:"column:code"`
    //名字
    Name string  `gorm:"column:name"`
    //类型   1对象 2对象list  3参数 5结果 9函数
    Type int  `gorm:"column:type"`
    //
    Class string  `gorm:"column:class"`
    //
    Db string  `gorm:"column:db"`
    //
    FormualSql string  `gorm:"column:formual_sql"`
    //匹配key
    Mapkey string  `gorm:"column:mapkey"`
    
}

//产品表名
func (IhrGruleObject) TableName() string {
	return "ihr_grule_object"
}
