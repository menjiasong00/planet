package service

import (
	"fmt"
	"golang.org/x/net/context"
	pb "planet/pb"
	"planet/pkg/gcode"
)


type TestServer struct{}

func (m *TestServer) GetTestMsg(c context.Context, s *pb.TestMessage) (*pb.TestMessage, error) {
	fmt.Printf("xxxxx(%q)\n", s.Value)
	//RunRulesResult(1," 1=1  ")
	gcode.MakeCoding(gcode.MakeCodingRequest{
		DatabaseName:"app_pro",
		TableName:"pro_products",
		Name:"商品",
		ServerName:"Pro",
		ModuleName:"Product",
		Env: "Pro",
	})
	return s, nil
}

