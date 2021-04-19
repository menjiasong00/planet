package model



//产品
type IhrGruleFormual struct {
    //
    Id int  `gorm:"column:id"`
    //规则集id
    GruleId int  `gorm:"column:grule_id"`
    //
    Name string  `gorm:"column:name"`
    //名字
    Info string  `gorm:"column:info"`
    //循环规则
    ForObject string  `gorm:"column:for_object"`
    //规则类别
    Class string  `gorm:"column:class"`
    //类型 1 正常规则 2循环规则
    Type int  `gorm:"column:type"`
    //排序
    Sort int  `gorm:"column:sort"`
    //#规则配置
    Formual string  `gorm:"column:formual"`
    //
    FormualText string  `gorm:"column:formual_text"`
    //
    Version string  `gorm:"column:version"`
    //
    Status int  `gorm:"column:status"`
    //
    IsDel int  `gorm:"column:is_del"`
    
}

//产品表名
func (IhrGruleFormual) TableName() string {
	return "ihr_grule_formual"
}
