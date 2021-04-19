package model



//产品
type IhrGrule struct {
    //
    Id int  `gorm:"column:id"`
    //名字
    Name string  `gorm:"column:name"`
    //
    Class string  `gorm:"column:class"`
    //
    ClassName string  `gorm:"column:class_name"`
    //
    Status int  `gorm:"column:status"`
    //
    Db string  `gorm:"column:db"`
    //入参
    InSql string  `gorm:"column:in_sql"`
    //结果输出
    OutTable string  `gorm:"column:out_table"`
    //结果选择字段
    OutColumns string  `gorm:"column:out_columns"`
    
}

//产品表名
func (IhrGrule) TableName() string {
	return "ihr_grule"
}
