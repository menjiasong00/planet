package service

import (
	"golang.org/x/net/context"

	db "planet/env/db"
	"planet/model"
	"planet/pb"
	"planet/pkg/tools"
)

type UsrServer struct{}

//UserList 用户列表
func (s *UsrServer) UserList(ctx context.Context, in *pb.UserListRequest) (*pb.UserListResponse, error) {

	resp := &pb.UserList{}

	dbModel := db.New("Usr").Model(model.UsrUser{})

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

	return &pb.UserListResponse{Status: 200, Message: "success", Details: resp}, nil
}

//UserDetail 用户详情
func (s *UsrServer) UserDetail(ctx context.Context, in *pb.UserIdRequest) (*pb.UserDetailResponse, error) {

	resp := &pb.UserOneRequest{}
	db.New("Usr").Model(model.UsrUser{}).Where("id = ?",in.Id).Scan(&resp)

	return &pb.UserDetailResponse{Status: 200, Message: "success", Details: resp}, nil
}


//UserCreate 用户新建
func (s *UsrServer) UserCreate(ctx context.Context, in *pb.UserOneRequest) (*pb.UserResponse, error) {
    
	//表单验证
    errValidate := in.Validate()
    	if errValidate != nil {
    		return nil, errValidate
    }

	UserOne := model.UsrUser{}
	tools.ScanStuct(in,&UserOne)

	db.New("Usr").Create(&UserOne)

	return &pb.UserResponse{Status: 200, Message: "success", Details:true}, nil
}

//UserMotify 用户修改
func (s *UsrServer) UserMotify(ctx context.Context, in *pb.UserOneRequest) (*pb.UserResponse, error) {
    
	//表单验证
    errValidate := in.Validate()
    	if errValidate != nil {
    		return nil, errValidate
    }

	UserOne := model.UsrUser{}
	db.New("Usr").Model(model.UsrUser{}).Where("id = ?",in.Id).Find(&UserOne)
	tools.ScanStuct(in,&UserOne)

	db.New("Usr").Model(model.UsrUser{}).Where(" id = ?",in.Id).Save(&UserOne)

	return &pb.UserResponse{Status: 200, Message: "success", Details:true}, nil
}

//UserDelete 用户删除
func (s *UsrServer) UserDelete(ctx context.Context, in *pb.UserIdRequest) (*pb.UserResponse, error) {

    UserOne := model.UsrUser{}
    db.New("Usr").Model(model.UsrUser{}).First(&UserOne,in.Id)
    //UserOne.Status = 2
    db.New("Usr").Save(&UserOne)
	
	return &pb.UserResponse{Status: 200, Message: "success", Details:true}, nil
}






