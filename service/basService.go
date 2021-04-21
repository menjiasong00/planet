package service

import (
	db "planet/env/db"
	"planet/model"
	"planet/pkg/gmysql"
	"planet/pkg/grule/engine"
	"planet/pkg/grule/fuc"
	"planet/pkg/tools"
	"regexp"
	"strings"
)

type BasServer struct{}



// DlxConsumer 死信消费者
/*

CREATE TABLE `bas_mq_dlx` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `exchange` varchar(255) NOT NULL COMMENT '交换机',
  `routing_key` json NOT NULL COMMENT '路由标识',
  `queue` varchar(255) NOT NULL COMMENT '队列名称',
  `app_id` varchar(255) NOT NULL COMMENT 'app_id',
  `body` json DEFAULT NULL COMMENT '消息体',
  `header` json DEFAULT NULL COMMENT '消息头',
  `expired_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '过期时间',
  `create_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
  `create_by` varchar(200) DEFAULT '',
  `update_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `update_by` varchar(200) DEFAULT '更新人',
  `status` tinyint(4) DEFAULT '1' COMMENT '## 1 未处理  2 已处理',
  PRIMARY KEY (`id`),
  KEY `EX_IDX` (`exchange`),
  KEY `QUEUE_IDX` (`queue`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='mq消息死信表';

message DlxConsumerRequest{
    string header =1;
    string body =2;
}

func (s *TestServer) DlxConsumer(ctx context.Context, in *pb.DlxConsumerRequest) (*pb.BaseResponse, error) {

	type Xheader struct {
		Count int `json:"count"`
		Exchange string `json:"exchange"`
		Queue string `json:"queue"`
		Reason string `json:"reason"`
		RoutingKeys []string `json:"routing-keys"`
		Time time.Time `json:"time"`
	}

	type Header struct {
		AppId string `json:"appid"`
		XDeath []Xheader `json:"x-death"`
		XFirstDeathExchange string `json:"x-first-death-exchange"`
		XFirstDeathQueue string `json:"x-first-death-queue"`
		XFirstDeathReason string `json:"x-first-death-reason"`
	}

	var dlxHeader Header
	json.Unmarshal([]byte(in.Header), &dlxHeader)

	var firstXheader Xheader

	for k,v:= range dlxHeader.XDeath{
		if (k==0){
			firstXheader = v
		}else{
			if v.Time.Unix() < firstXheader.Time.Unix() {
				firstXheader = v
			}
		}
	}
	if firstXheader.Exchange =="" {
		firstXheader.Time = time.Now()
	}

	routingkey ,_:= json.Marshal(firstXheader.RoutingKeys)
	newBasMqDlx := model.BasMqDlx{
		Exchange:firstXheader.Exchange,
		Queue:firstXheader.Queue,
		Header:in.Header,
		Body:in.Body,
		CreateAt:firstXheader.Time,
		UpdateAt:firstXheader.Time,
		ExpiredAt:firstXheader.Time,
		Status:1,
		AppId:dlxHeader.AppId,
		RoutingKey: bytes.NewBuffer(routingkey).String(),
	}


	err:= db.New("Bas").Save(&newBasMqDlx).Error

	if err !=nil {
		return &pb.BaseResponse{}, err
	}

	return &pb.BaseResponse{}, nil
}
*/

 
func RunRulesResult(gruleId int,inWhere string) error{

	//入参列表
	var listGrule model.IhrGrule
	db.New("Bas").Where("`id` = ?",gruleId).First(&listGrule)
	inSql:= listGrule.InSql +" AND  "+ inWhere
	mapList, err2 := gmysql.MapByQuery(db.New("Bas").DB(), inSql, "string")
	//fmt.Println(inSql)
	if err2!=nil {
		return err2
	}

	//对象配置
	var listGruleObjects []model.IhrGruleObject
	db.New("Bas").Where("(`class` = ?  or code = 'person') and type in (1,2)",listGrule.Class).Find(&listGruleObjects)

	//变量配置
	var listGruleItems []model.IhrGruleItem
	db.New("Bas").Table("ihr_grule_item").Where("  (`class` = ? or type=9 or object= 'person')   AND status = 1",listGrule.Class).Order("`type`asc,sort asc").Find(&listGruleItems)

	//规则配置
	var listGruleFormual []model.IhrGruleFormual
	db.New("Bas").Table("ihr_grule_formual").Where("`grule_id` = ? AND status = 1",listGrule.Id).Order("sort asc").Find(&listGruleFormual)

	poolMap:= make(map[int]*engine.GenginePool,0 )


	for _,vItem := range listGruleFormual {
		rulesString,err :=MakeRulesCode(listGruleItems,vItem)
		if err != nil{
			return err
		}

		//普通规则
		rulesFile:=`
rule "rule name" "rule desc"
begin
  `+rulesString+`
end
`
		var Fuc fuc.BasFunc
		apis := make(map[string]interface{})
		//apis["PrintName"] = fmt.Println
		apis["Fuc"] = &Fuc
		pool, e0 := engine.NewGenginePool(1, 8, 1, rulesFile, apis)
		if e0 != nil{
			return e0
		}
		poolMap[vItem.Id] = pool
	}


	var insertList []map[string]interface{}
	//初始化 对象属性

	for _,v:= range mapList {
		//fmt.Println(v)
		mapObjList:= make (map[string]map[string]interface{},0)
		mapObjsList:= make (map[string][]map[string]interface{},0)

		//初始化 对象属性
		for _,vo:= range listGruleObjects{
			for kp,vp := range v{
				vo.FormualSql = strings.Replace(vo.FormualSql,"{"+kp+"}","'"+tools.Strval(vp)+"'",-1)
			}
			mapList, err2 := gmysql.MapByQuery(db.New("Bas").DB(), vo.FormualSql, "string")
			if err2 != nil {
			}
			if len(mapList) > 0  && vo.Type == 1 {
				mapObjList[vo.Code] = mapList[0]
			}
			if len(mapList) > 0  && vo.Type == 2 {
				mapObjsList[vo.Code] = mapList
			}
		}

		resultMap := make(map[string]interface{},0)

		mapCodeListGrule := make(map[string]model.IhrGruleItem,0)
		mapNameListGrule := make(map[string]model.IhrGruleItem,0)
		mapNameFuncFormual := make(map[string]string,0)

		StringsMap:= make(map[string]string,0)
		IntsMap:= make(map[string]int,0)
		Float64sMap:= make(map[string]float64,0)


		for _,v:= range listGruleItems{
			key:= "【"+v.ObjectName+"."+v.Name+"】"
			mapNameListGrule[key] = v
			mapCodeListGrule[v.Object+"."+v.Code] = v
			if v.Type ==9 {
				mapNameFuncFormual[key] = v.FormualShow
			}else{
				key :=  v.Object+"."+v.Code
				if v.FiledType == "string" {
					StringsMap[key] = ""
				}else if v.FiledType == "int" {
					IntsMap[key] = 0
				}else if v.FiledType == "float64" {
					Float64sMap[key] = 0
				}
			}
		}

		objList :=  mapObjList
		objsList :=  mapObjsList

		for kobj,vobj := range objList {
			for k,v := range vobj {
				if aObj, ok := mapCodeListGrule[kobj+"."+k]; ok {
					key :=  aObj.Object+"."+aObj.Code
					if aObj.FiledType == "string" {
						StringsMap[key] = v.(string)
					}else if aObj.FiledType == "int" {
						IntsMap[key] = tools.StrToInt(v.(string))
					}else if aObj.FiledType == "float64" {
						Float64sMap[key] = tools.StrToFlout(v.(string))
					}
				}
			}
		}

		Data := &ParamsData{
			Strings:StringsMap,
			Ints:IntsMap,
			Float64s:Float64sMap,
		}

		for _,vItem := range listGruleFormual {
			//普通规则
			poolOne:=  poolMap[vItem.Id]
			if vItem.Type ==1 {
				e2 ,_:= poolOne.ExecuteRulesWithSpecifiedEM("Data", Data, "", nil)
				if e2 != nil{
					return e2
				}
			}else if  vItem.Type ==2 {
				//循环规则，替换循环对象
				for kobjs,vobjs := range objsList {
					if kobjs == vItem.ForObject {
						//循环规则，执行对应对象的规则
						for _,vobj := range vobjs {
							for k,v := range vobj {

								if aObj, ok := mapCodeListGrule[kobjs+"."+k]; ok {
									key :=  aObj.Object+"."+aObj.Code

									//key :=  aObj.Object+tools.Strval(kobj)+"."+aObj.Code
									if aObj.FiledType == "string" {
										StringsMap[key] = v.(string)
									}else if aObj.FiledType == "int" {
										IntsMap[key] = tools.StrToInt(v.(string))
									}else if aObj.FiledType == "float64" {
										Float64sMap[key] = tools.StrToFlout(v.(string))
									}
								}
							}

							//执行
							Data.Strings = StringsMap
							Data.Ints = IntsMap
							Data.Float64s = Float64sMap

							e2 ,_:= poolOne.ExecuteRulesWithSpecifiedEM("Data", Data, "", nil)
							if e2 != nil{
								return e2
							}
						}
					}
				}
			}

		}



		for k,v:= range Data.Strings {
			resultMap[k] = v
		}
		for k,v:= range Data.Ints {
			resultMap[k] = v
		}
		for k,v:= range Data.Float64s {
			resultMap[k] = v
		}


		oneOutMap := make(map[string]interface{})
		for kr,vr := range resultMap{
			nameArr := strings.Split(kr, ".")
			oneOutMap[nameArr[1]] = vr
		}

		insertList = append(insertList,oneOutMap)

	}


	err3 := gmysql.ReplaceInsertMany(db.New("Bas").DB(),listGrule.OutTable,insertList)
	if err3!= nil {
		return err3
	}
	return nil

}
type ParamsData struct {
	Strings map[string]string
	Ints map[string]int
	Float64s map[string]float64
	Bools  map[string]bool
}


func MakeRulesCode (listGruleItems []model.IhrGruleItem,gruleFormual  model.IhrGruleFormual) (string,error){
	mapCodeListGrule := make(map[string]model.IhrGruleItem,0)
	mapNameListGrule := make(map[string]model.IhrGruleItem,0)
	mapNameFuncFormual := make(map[string]string,0)
	for _,v:= range listGruleItems{
		key:= "【"+v.ObjectName+"."+v.Name+"】"
		mapNameListGrule[key] = v
		mapCodeListGrule[v.Object+"."+v.Code] = v
		if v.Type ==9 {
			mapNameFuncFormual[key] = v.FormualShow
		}
	}

	keyMap := []string{"【如果】","【那么】","【结束】","【否则如果】","【否则】","【或】","【且】"}

	rulesString :=  gruleFormual.Formual
	rulesString  =  strings.Replace(rulesString,"【如果】"," if ",-1)
	rulesString  =  strings.Replace(rulesString,"【那么】"," { ",-1)
	rulesString  =  strings.Replace(rulesString,"【结束】"," } ",-1)
	rulesString  =  strings.Replace(rulesString,"【否则如果】","} else if ",-1)
	rulesString  =  strings.Replace(rulesString,"【否则】"," } else { ",-1)
	rulesString  =  strings.Replace(rulesString,"【或】"," || ",-1)
	rulesString  =  strings.Replace(rulesString,"【且】"," && ",-1)
	//普通规则
	exp2 := regexp.MustCompile(`【(.*?)】`)
	result2 := exp2.FindAllStringSubmatch(gruleFormual.Formual, -1)
	if len(result2) > 0 {

		for _, vr := range result2 {
			if aObj, ok := mapNameListGrule[vr[0]]; ok {
				if aObj.Type ==9 {
					//公式
					rulesString  =  strings.Replace(rulesString,vr[0],aObj.FormualShow,-1)
				}else{
					key :=  aObj.Object+"."+aObj.Code

					if aObj.FiledType == "string" {
						rulesString  =  strings.Replace(rulesString,vr[0],`Data.Strings["`+key+`"]`,-1)
					}else if aObj.FiledType == "int" {
						rulesString  =  strings.Replace(rulesString,vr[0],`Data.Ints["`+key+`"]`,-1)
					}else if aObj.FiledType == "float64" {
						rulesString  =  strings.Replace(rulesString,vr[0],`Data.Float64s["`+key+`"]`,-1)
					}
				}
			}else{
				if tools.StrInIface(vr[0],keyMap) == false {
					//println("没有匹配到"+vr[0])
					//return  "", errors.New("没有匹配到关键字："+vr[0])
				}
			}
		}
	}

	return rulesString,nil

}

