package model



//产品
type IhrGruleItem struct {
    //
    Id int  `gorm:"column:id"`
    //
    Code string  `gorm:"column:code"`
    //名字
    Name string  `gorm:"column:name"`
    //
    Object string  `gorm:"column:object"`
    //
    ObjectName string  `gorm:"column:object_name"`
    //
    Class string  `gorm:"column:class"`
    //类型 1中间参数 2基础表属性  5结果 9函数
    Type int  `gorm:"column:type"`
    //排序
    Sort int  `gorm:"column:sort"`
    //#规则配置
    FormualShow string  `gorm:"column:formual_show"`
    //类型
    FiledType string  `gorm:"column:filed_type"`
    //
    Version string  `gorm:"column:version"`
    //
    Status int  `gorm:"column:status"`
    
}

//产品表名
func (IhrGruleItem) TableName() string {
	return "ihr_grule_item"
}
