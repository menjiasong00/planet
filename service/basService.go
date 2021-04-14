package service

import (
	"golang.org/x/net/context"
	"planet/pkg/gmysql"
	"planet/model"
	"planet/pb"
	"planet/pkg/tools"
)

type BasServer struct{}

//BasProductsList 产品列表
func (s *BasServer) BasProductsList(ctx context.Context, in *pb.BasProductsListRequest) (*pb.BasProductsListResponse, error) {

	resp := &pb.BasProductsList{}

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

	return &pb.BasProductsListResponse{Status: 200, Message: "success", Details: resp}, nil
}

//BasProductsDetail 产品详情
func (s *BasServer) BasProductsDetail(ctx context.Context, in *pb.BasProductsIdRequest) (*pb.BasProductsDetailResponse, error) {

	resp := &pb.BasProductsOneRequest{}
	gmysql.DB.Model(model.Products{}).Where("id = ?",in.Id).Scan(&resp)

	return &pb.BasProductsDetailResponse{Status: 200, Message: "success", Details: resp}, nil
}


//BasProductsCreate 产品新建
func (s *BasServer) BasProductsCreate(ctx context.Context, in *pb.BasProductsOneRequest) (*pb.BasProductsResponse, error) {
    
	//表单验证
    errValidate := in.Validate()
    	if errValidate != nil {
    		return nil, errValidate
    }

	BasProductsOne := model.Products{}
	tools.ScanStuct(in,&BasProductsOne)

	gmysql.DB.Create(&BasProductsOne)

	return &pb.BasProductsResponse{Status: 200, Message: "success", Details:true}, nil
}

//BasProductsMotify 产品修改
func (s *BasServer) BasProductsMotify(ctx context.Context, in *pb.BasProductsOneRequest) (*pb.BasProductsResponse, error) {
    
	//表单验证
    errValidate := in.Validate()
    	if errValidate != nil {
    		return nil, errValidate
    }

	BasProductsOne := model.Products{}
	gmysql.DB.Model(model.Products{}).Where("id = ?",in.Id).Find(&BasProductsOne)
	tools.ScanStuct(in,&BasProductsOne)

	gmysql.DB.Model(model.Products{}).Where(" id = ?",in.Id).Save(&BasProductsOne)

	return &pb.BasProductsResponse{Status: 200, Message: "success", Details:true}, nil
}

//BasProductsDelete 产品删除
func (s *BasServer) BasProductsDelete(ctx context.Context, in *pb.BasProductsIdRequest) (*pb.BasProductsResponse, error) {

    BasProductsOne := model.Products{}
    gmysql.DB.Model(model.Products{}).First(&BasProductsOne,in.Id)
    //BasProductsOne.Status = 2
    gmysql.DB.Save(&BasProductsOne)
	
	return &pb.BasProductsResponse{Status: 200, Message: "success", Details:true}, nil
}












