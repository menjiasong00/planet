package service

import (
	"golang.org/x/net/context"
	"planet/pkg/gmysql"
	"planet/model"
	"planet/pb"
	"planet/pkg/tools"
)

type BasServer struct{}

//BaslProductsList 产品列表
func (s *BasServer) BaslProductsList(ctx context.Context, in *pb.BaslProductsListRequest) (*pb.BaslProductsListResponse, error) {

	resp := &pb.BaslProductsList{}

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

	return &pb.BaslProductsListResponse{Status: 200, Message: "success", Details: resp}, nil
}

//BaslProductsDetail 产品详情
func (s *BasServer) BaslProductsDetail(ctx context.Context, in *pb.BaslProductsIdRequest) (*pb.BaslProductsDetailResponse, error) {

	resp := &pb.BaslProductsOneRequest{}
	gmysql.DB.Model(model.Products{}).Where("id = ?",in.Id).Scan(&resp)

	return &pb.BaslProductsDetailResponse{Status: 200, Message: "success", Details: resp}, nil
}


//BaslProductsCreate 产品新建
func (s *BasServer) BaslProductsCreate(ctx context.Context, in *pb.BaslProductsOneRequest) (*pb.BaslProductsResponse, error) {
    
	//表单验证
    errValidate := in.Validate()
    	if errValidate != nil {
    		return nil, errValidate
    }

	BaslProductsOne := model.Products{}
	tools.ScanStuct(in,&BaslProductsOne)

	gmysql.DB.Create(&BaslProductsOne)

	return &pb.BaslProductsResponse{Status: 200, Message: "success", Details:true}, nil
}

//BaslProductsMotify 产品修改
func (s *BasServer) BaslProductsMotify(ctx context.Context, in *pb.BaslProductsOneRequest) (*pb.BaslProductsResponse, error) {
    
	//表单验证
    errValidate := in.Validate()
    	if errValidate != nil {
    		return nil, errValidate
    }

	BaslProductsOne := model.Products{}
	gmysql.DB.Model(model.Products{}).Where("id = ?",in.Id).Find(&BaslProductsOne)
	tools.ScanStuct(in,&BaslProductsOne)

	gmysql.DB.Model(model.Products{}).Where(" id = ?",in.Id).Save(&BaslProductsOne)

	return &pb.BaslProductsResponse{Status: 200, Message: "success", Details:true}, nil
}

//BaslProductsDelete 产品删除
func (s *BasServer) BaslProductsDelete(ctx context.Context, in *pb.BaslProductsIdRequest) (*pb.BaslProductsResponse, error) {

    BaslProductsOne := model.Products{}
    gmysql.DB.Model(model.Products{}).First(&BaslProductsOne,in.Id)
    //BaslProductsOne.Status = 2
    gmysql.DB.Save(&BaslProductsOne)
	
	return &pb.BaslProductsResponse{Status: 200, Message: "success", Details:true}, nil
}






