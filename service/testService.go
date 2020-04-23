package service

import(
	pb "planet/pb"
	"golang.org/x/net/context"
	"fmt"
)


type TestServer struct{}
func (m *TestServer) GetTestMsg(c context.Context, s *pb.TestReq) (*pb.TestResp, error) {
	fmt.Printf("xxxxx(%q)\n", s.Value)
	details := &pb.TestDetail{Value:s.Value}
	return &pb.TestResp{Details:details}, nil
}



//自动编码
/* func (s *BasMqServer) MakeCoding(ctx context.Context, in *pb.MakeCodingRequest) (*pb.DlxResendResponse, error) {

	//sweaters := Inventory{"BasInfo", "Item", "BasInfoClassItemT","basInfo", "item"}

	sweaters := Inventory{in.Name,in.ServerName, in.ModuleName,"" ,"", "",in.TableName,nil}

	moduleNameSnake := tools.SnakeString(sweaters.ModuleName)
	serverNameSnake := tools.SnakeString(sweaters.ServerName)
	sweaters.ModuleNameLow = tools.SnakeToCamel(moduleNameSnake)
	sweaters.ServerNameLow = tools.SnakeToCamel(serverNameSnake)
	sweaters.ModelName = tools.SnakeToBigCamel(sweaters.TableName)
	//SELECT COLUMN_NAME as column_name, column_comment FROM INFORMATION_SCHEMA. COLUMNS WHERE table_name = 'bas_mq_dlx' AND table_schema = 'wos_common'

	var result []SqlResult
	// Raw SQL
	db.SysDB.Raw("SELECT COLUMN_NAME as column_name,data_type, column_comment FROM INFORMATION_SCHEMA. COLUMNS WHERE table_name = ? AND table_schema = ?", sweaters.TableName,in.DatabaseName).Scan(&result)
	i := 1
	for k,v:= range result {
		result[k].ColumnNameCamel = tools.SnakeToCamel(v.ColumnName)
		result[k].ColumnNameBigCamel = tools.SnakeToBigCamel(v.ColumnName)
		result[k].Id = i
		i++

		switch v.DataType {
			case "int","tinyint","smallint","mediumint":
				result[k].StuctTypeName = "int"
				result[k].TypeName = "int32"
			case "varchar","char","text","tinytext","mediumtext","longtext":
				result[k].StuctTypeName = "string"
				result[k].TypeName = "string"
			case "double":
				result[k].StuctTypeName = "float32"
				result[k].TypeName = "float"
			case "bool":
				result[k].StuctTypeName = "bool"
				result[k].TypeName = "bool"
			case "float":
				result[k].StuctTypeName = "float64"
				result[k].TypeName = "float"
			case "timestamp","date","datetime","time":
				result[k].StuctTypeName = "time.Time"
				result[k].TypeName = "string"
			default:
				result[k].StuctTypeName = "string"
				result[k].TypeName = "string"
		//
		}
	}

	sweaters.Columns = result

	MakeFile(sweaters,"./pkg/coding/model.go.tpl","./model/"+sweaters.TableName+".go-")
	MakeFile(sweaters,"./pkg/coding/proto.go.tpl","./proto/"+sweaters.ServerNameLow+".proto-")
	MakeFile(sweaters,"./pkg/coding/server.go.tpl","./service/"+sweaters.ServerNameLow+"Service.go-")

	return &pb.DlxResendResponse{Status: 200, Message: "success", Data:true}, nil
}

type SqlResult struct {
	Id int
	DataType string
	ColumnName string
	ColumnComment string
	TypeName string
	ColumnNameCamel string
	ColumnNameBigCamel string
	StuctTypeName string
}

type Inventory struct {
	Name string
	ServerName string
	ModuleName string
	ModelName string
	ServerNameLow string
	ModuleNameLow string
	TableName string
	Columns []SqlResult
}

func MakeFile(sweaters Inventory,tmplFile string,outFile string) {

	tmpl, _ := template.ParseFiles(tmplFile)

	_ = tmpl.Execute(os.Stdout, sweaters)

	f, _ := os.OpenFile(outFile, os.O_WRONLY|os.O_CREATE, 0644)
	defer f.Close()

	// 渲染并写入文件
	_ = tmpl.Execute(f, sweaters)

} */