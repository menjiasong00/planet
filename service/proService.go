package service

import (
	"golang.org/x/net/context"

	"planet/model"
	"planet/pb"
	"planet/pkg/tools"
	db "planet/env/db"
)

type ProServer struct{}

//ProductList 商品列表
func (s *ProServer) ProductList(ctx context.Context, in *pb.ProductListRequest) (*pb.ProductListResponse, error) {

	resp := &pb.ProductList{}

	dbModel := db.New("Pro").Model(model.ProProducts{})

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

	return &pb.ProductListResponse{Status: 200, Message: "success", Details: resp}, nil
}

//ProductDetail 商品详情
func (s *ProServer) ProductDetail(ctx context.Context, in *pb.ProductIdRequest) (*pb.ProductDetailResponse, error) {

	resp := &pb.ProductOneRequest{}
	db.New("Pro").Model(model.ProProducts{}).Where("id = ?",in.Id).Scan(&resp)

	return &pb.ProductDetailResponse{Status: 200, Message: "success", Details: resp}, nil
}


//ProductCreate 商品新建
func (s *ProServer) ProductCreate(ctx context.Context, in *pb.ProductOneRequest) (*pb.ProductResponse, error) {
    
	//表单验证
    errValidate := in.Validate()
    	if errValidate != nil {
    		return nil, errValidate
    }

	ProductOne := model.ProProducts{}
	tools.ScanStuct(in,&ProductOne)

	db.New("Pro").Create(&ProductOne)

	return &pb.ProductResponse{Status: 200, Message: "success", Details:true}, nil
}

//ProductMotify 商品修改
func (s *ProServer) ProductMotify(ctx context.Context, in *pb.ProductOneRequest) (*pb.ProductResponse, error) {
    
	//表单验证
    errValidate := in.Validate()
    	if errValidate != nil {
    		return nil, errValidate
    }

	ProductOne := model.ProProducts{}
	db.New("Pro").Model(model.ProProducts{}).Where("id = ?",in.Id).Find(&ProductOne)
	tools.ScanStuct(in,&ProductOne)

	db.New("Pro").Model(model.ProProducts{}).Where(" id = ?",in.Id).Save(&ProductOne)

	return &pb.ProductResponse{Status: 200, Message: "success", Details:true}, nil
}

//ProductDelete 商品删除
func (s *ProServer) ProductDelete(ctx context.Context, in *pb.ProductIdRequest) (*pb.ProductResponse, error) {

    ProductOne := model.ProProducts{}
    db.New("Pro").Model(model.ProProducts{}).First(&ProductOne,in.Id)
    //ProductOne.Status = 2
    db.New("Pro").Save(&ProductOne)
	
	return &pb.ProductResponse{Status: 200, Message: "success", Details:true}, nil
}






