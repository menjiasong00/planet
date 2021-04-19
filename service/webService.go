package service

import (
	"golang.org/x/net/context"
	"planet/pkg/gmysql"
	"planet/model"
	"planet/pb"
	"planet/pkg/tools"
)

type WebServer struct{}

//ProductsList 产品列表
func (s *WebServer) ProductsList(ctx context.Context, in *pb.ProductsListRequest) (*pb.ProductsListResponse, error) {

	resp := &pb.ProductsList{}

	dbModel := gmysql.DB.Model(model.Products{})

/*	if len(in.Name) > 0 {
		dbModel=dbModel.Where("  name like ?","%"+in.Name+"%")
	}*/

	dbModel.Count(&resp.Total)

	if in.PageSize > 0 {
		//分页
		offset := (in.PageNumber - 1) * in.PageSize
		dbModel = dbModel.Offset(offset).Limit(in.PageSize)
	}
	
	sortStr := " id DESC"
	if in.OrderKey != "" && in.OrderSort != "" {
		sortStr = in.OrderKey + " " + in.OrderSort
	}
	dbModel = dbModel.Order(sortStr)

	dbModel.Scan(&resp.List)

/*	for _, v := range resp.List {

	}*/

	return &pb.ProductsListResponse{Status: 200, Message: "success", Details: resp}, nil
}

//ProductsDetail 产品详情
func (s *WebServer) ProductsDetail(ctx context.Context, in *pb.ProductsIdRequest) (*pb.ProductsDetailResponse, error) {

	resp := &pb.ProductsOneRequest{}
	gmysql.DB.Model(model.Products{}).Where("id = ?",in.Id).Scan(&resp)

	return &pb.ProductsDetailResponse{Status: 200, Message: "success", Details: resp}, nil
}


//ProductsCreate 产品新建
func (s *WebServer) ProductsCreate(ctx context.Context, in *pb.ProductsOneRequest) (*pb.ProductsResponse, error) {
    
	//表单验证
    errValidate := in.Validate()
    	if errValidate != nil {
    		return nil, errValidate
    }

	ProductsOne := model.Products{}
	tools.ScanStuct(in,&ProductsOne)

	gmysql.DB.Create(&ProductsOne)

	return &pb.ProductsResponse{Status: 200, Message: "success", Details:true}, nil
}

//ProductsMotify 产品修改
func (s *WebServer) ProductsMotify(ctx context.Context, in *pb.ProductsOneRequest) (*pb.ProductsResponse, error) {
    
	//表单验证
    errValidate := in.Validate()
    	if errValidate != nil {
    		return nil, errValidate
    }

	ProductsOne := model.Products{}
	gmysql.DB.Model(model.Products{}).Where("id = ?",in.Id).Find(&ProductsOne)
	tools.ScanStuct(in,&ProductsOne)

	gmysql.DB.Model(model.Products{}).Where(" id = ?",in.Id).Save(&ProductsOne)

	return &pb.ProductsResponse{Status: 200, Message: "success", Details:true}, nil
}

//ProductsDelete 产品删除
func (s *WebServer) ProductsDelete(ctx context.Context, in *pb.ProductsIdRequest) (*pb.ProductsResponse, error) {

    ProductsOne := model.Products{}
    gmysql.DB.Model(model.Products{}).First(&ProductsOne,in.Id)
    //ProductsOne.Status = 2
    gmysql.DB.Save(&ProductsOne)
	
	return &pb.ProductsResponse{Status: 200, Message: "success", Details:true}, nil
}






